from flask import Flask, request, jsonify, Response
from flask_cors import CORS
import json
import os

from db import DBOperator
from llm import ChatClient
from rpc import QuestionRpcClient

# ==================== 配置参数 ====================
AI_API_KEY = os.getenv("AI_API_KEY")
AI_BASE_URL = os.getenv("AI_BASE_URL", "https://api.deepseek.com")
AI_MODEL = os.getenv("AI_MODEL", "deepseek-chat")

MYSQL_USER = os.getenv("MYSQL_USER")
MYSQL_PASSWORD = os.getenv("MYSQL_PASSWORD")
MYSQL_HOST = os.getenv("MYSQL_HOST")
MYSQL_DATABASE = os.getenv("MYSQL_DATABASE")

DATABASE_URI = (
    f"mysql+pymysql://{MYSQL_USER}:{MYSQL_PASSWORD}@{MYSQL_HOST}/{MYSQL_DATABASE}?charset=utf8mb4"
)

# ==================== 初始化 ====================
db_operator = DBOperator(DATABASE_URI)
chat_client = ChatClient(AI_API_KEY, AI_BASE_URL, AI_MODEL)
question_rpc = QuestionRpcClient()
app = Flask(__name__)


# ==================== 代码诊断接口 ====================
@app.route("/code/check", methods=["POST"])
def code_check():
    data = request.get_json()
    code = data.get("code")
    question_id = data.get("question_id")
    question_info = question_rpc.get_question_info(question_id)
    return jsonify(
        {
            "code": 200,
            "msg": "success",
            "data": chat_client.code_check(code, question_info),
        }
    )


# ==================== 流式输出接口 ====================
@app.route("/chat", methods=["POST"])
def chat():
    data = request.get_json()
    user_id = data.get("user_id")
    question_id = data.get("question_id")
    session_id_input = data.get("session_id")
    user_message = data.get("message")

    if not (user_id and question_id and user_message):
        return jsonify({"code": 400, "msg": "缺少必要参数", "data": None})

    try:
        # 获取或创建会话（同时插入系统提示）
        session_id = db_operator.get_or_create_session(
            session_id_input, user_id, question_id
        )
        # 获取下一轮轮次
        current_round = db_operator.get_next_round(session_id)
        title = ""
        # 如果当前轮次为1，则生成对话标题
        if current_round == 1:
            title = chat_client.generate_chat_title(user_message)
            db_operator.update_session_title(session_id, title)
        # 插入用户消息
        db_operator.insert_message(session_id, current_round, "user", user_message)
        # 构建对话上下文
        messages = db_operator.build_message_context(session_id)
    except Exception as e:
        return jsonify({"code": 500, "msg": f"系统错误: {str(e)}", "data": None})

    def generate():
        ai_reply_full = ""
        try:
            # 调用 openai 库，启用流式输出
            response = chat_client.chat(messages)
            # 遍历流数据，每次返回 chunk
            id = 1
            for chunk in response:
                delta = chunk.choices[0].delta
                if delta.content is not None:
                    chunk_text = delta.content
                    print(chunk_text)
                    ai_reply_full += chunk_text
                    yield f"event: message\ndata: {json.dumps(
                        {
                            'id': id,
                            'content': chunk_text,
                        }
                    )}\n\n"
                    id += 1
            # 将完整回复保存到数据库（轮次递增）
            db_operator.insert_message(
                session_id, current_round + 1, "assistant", ai_reply_full
            )
            yield f"event: message\ndata: {json.dumps(
                {
                    'msg': 'DONE',
                    'full_content': ai_reply_full,
                    'session_id': session_id,
                    'title': title,
                }
            )}\n\n"
            # 添加结束事件
            yield f"event: end\ndata: [DONE]\n\n"
        except Exception as e:
            yield f"event: error\ndata: {json.dumps(
                { 'msg': f'AI调用错误: {str(e)}'}
            )}\n\n"
            # 即使出错也发送结束事件
            yield f"event: end\ndata: [DONE]\n\n"

    return Response(generate(), mimetype="text/event-stream")


# ==================== 获取一道题的所有会话 ====================
@app.route("/question/sessions", methods=["POST"])
def get_question_sessions():
    data = request.get_json()
    user_id = data.get("user_id")
    question_id = data.get("question_id")
    if not (user_id and question_id):
        return jsonify({"code": 400, "msg": "缺少必要参数", "data": None})
    try:
        return jsonify(
            {
                "code": 200,
                "msg": "success",
                "data": db_operator.get_question_sessions(user_id, question_id),
            }
        )
    except Exception as e:
        return jsonify({"code": 500, "msg": f"系统错误: {str(e)}", "data": None})


# ==================== 获取一个会话的对话历史 ====================
@app.route("/session/chat/history", methods=["POST"])
def get_session_message_history():
    data = request.get_json()
    user_id = data.get("user_id")
    session_id = data.get("session_id")
    if not (user_id and session_id):
        return jsonify({"code": 400, "msg": "缺少必要参数", "data": None})
    try:
        return jsonify(
            {
                "code": 200,
                "msg": "success",
                "data": db_operator.get_session_message_history(user_id, session_id),
            }
        )
    except Exception as e:
        return jsonify({"code": 500, "msg": f"系统错误: {str(e)}", "data": None})


# ==================== 启动服务 ====================
if __name__ == "__main__":
    CORS(app)
    app.run(host="0.0.0.0", port=6005)
