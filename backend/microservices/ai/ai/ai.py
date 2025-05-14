import openai


class ChatClient:
    def __init__(self, api_key: str, base_url: str, model: str = "deepseek-chat"):
        self.api_key = api_key
        self.base_url = base_url
        self.model = model
        self.client = openai.OpenAI(api_key=self.api_key, base_url=self.base_url)

    # 对用户上传的代码进行代码问题诊断检测代码存在的错误
    def code_check(self, code: str, question_info: dict) -> str:
        prompt = f"""作为代码检查专家，请用简洁易懂的语言指出以下代码的主要问题，用1-3句话说明关键错误和改进建议：
                题目：\n{question_info["title"]}
                题目内容（markdown格式）：\n{question_info["content"]}
                题目标签（json格式）：\n{question_info["tags"]}
                代码：\n{code}
                请直接指出代码最需要修正的1-2个问题，并用自然的口语化表达给出具体改进建议。"""
        print(prompt)
        response = self.client.chat.completions.create(
            model=self.model,
            messages=[
                {
                    "role": "system",
                    "content": "你是一个专业的代码检查助手，善于诊断代码存在的问题。",
                },
                {"role": "user", "content": prompt},
            ],
            stream=False,
        )
        return response.choices[0].message.content

    def chat(self, messages: list):
        response = self.client.chat.completions.create(
            model=self.model,
            messages=messages,
            stream=True,
        )
        return response

    def generate_chat_title(self, first_user_message: str) -> str:
        prompt = f"""你是一个智能助手，善于提炼对话主题。
        请根据用户的第一句消息，生成一个简洁、准确的对话标题，适合用作对话列表中的名称。
        要求：
        - 不超过10个字
        - 准确概括核心主题
        - 不要加标点或引号
        - 使用简体中文

        示例：
        用户：我想学怎么用Go调用MySQL数据库  
        标题：Go操作MySQL

        用户：请问如何提高英语口语？  
        标题：英语口语提升

        用户：{first_user_message}  
        标题："""

        response = self.client.chat.completions.create(
            model=self.model,
            messages=[
                {"role": "system", "content": "你是一个善于总结主题的助手。"},
                {"role": "user", "content": prompt},
            ],
            stream=False,
        )

        title = response.choices[0].message.content.strip()
        return title
