# db/ops.py

import uuid
from sqlalchemy import (
    Column,
    BigInteger,
    String,
    Text,
    Enum,
    TIMESTAMP,
    func,
    create_engine,
)
from sqlalchemy.orm import sessionmaker, scoped_session, Session as SessionType
from sqlalchemy.ext.declarative import declarative_base

from rpc.rpc import QuestionRpcClient

DB_USER = ""
DB_PASSWORD = ""
DB_HOST = ""
DB_NAME = ""
DATABASE_URI = (
    f"mysql+pymysql://{DB_USER}:{DB_PASSWORD}@{DB_HOST}/{DB_NAME}?charset=utf8mb4"
)

Base = declarative_base()


class SessionModel(Base):
    __tablename__ = "sessions"
    __table_args__ = {"comment": "对话会话表"}

    id = Column(BigInteger, primary_key=True, autoincrement=True)
    session_id = Column(String(64), nullable=False, unique=True)
    title = Column(String(255), nullable=False)
    user_id = Column(BigInteger, nullable=False)
    question_id = Column(BigInteger, nullable=False)
    create_at = Column(TIMESTAMP, nullable=False, server_default=func.now())
    update_at = Column(
        TIMESTAMP, nullable=False, server_default=func.now(), onupdate=func.now()
    )


class MessageModel(Base):
    __tablename__ = "messages"
    __table_args__ = {"comment": "对话消息表"}

    id = Column(BigInteger, primary_key=True, autoincrement=True)
    session_id = Column(String(64), nullable=False)
    round = Column(BigInteger, nullable=False)
    message_role = Column(Enum("system", "assistant", "user"), nullable=False)
    content = Column(Text, nullable=False)
    create_at = Column(TIMESTAMP, nullable=False, server_default=func.now())
    update_at = Column(
        TIMESTAMP, nullable=False, server_default=func.now(), onupdate=func.now()
    )


class DBManager:
    def __init__(self):
        self.engine = create_engine(DATABASE_URI, pool_pre_ping=True)
        self.Session = scoped_session(sessionmaker(bind=self.engine))
        self.question_rpc = QuestionRpcClient()

    def build_system_prompt(self, question_text: str) -> dict:
        content = (
            f"你是一个编程竞赛助手，当前的问题是：\n{question_text}\n"
            "请根据用户的问题提供帮助，回答要简洁，专注于算法思路"
        )
        return {"role": "system", "content": content}

    def get_session(self) -> SessionType:
        return self.Session()

    def get_or_create_session(self, session_id: str, user_id: int, question_id: int):
        db = self.get_session()
        try:
            if session_id:
                session_obj = (
                    db.query(SessionModel).filter_by(session_id=session_id).first()
                )
                if session_obj:
                    return session_obj.session_id

            new_session_id = str(uuid.uuid4())
            session_obj = SessionModel(
                session_id=new_session_id,
                user_id=user_id,
                question_id=question_id,
                title="暂无标题",
            )
            db.add(session_obj)
            db.commit()

            # 添加系统提示词
            question_info = self.question_rpc.get_question_info(question_id)
            system_msg = self.build_system_prompt(question_info["content"])
            msg = MessageModel(
                session_id=new_session_id,
                round=0,
                message_role=system_msg["role"],
                content=system_msg["content"],
            )
            db.add(msg)
            db.commit()

            return session_obj.session_id
        finally:
            db.close()

    def build_message_context(self, session_id: str):
        db = self.get_session()
        try:
            records = (
                db.query(MessageModel)
                .filter_by(session_id=session_id)
                .order_by(MessageModel.round)
                .all()
            )
            return [
                {"role": rec.message_role, "content": rec.content} for rec in records
            ]
        finally:
            db.close()

    def insert_message(self, session_id: str, round_num: int, role: str, content: str):
        db = self.get_session()
        try:
            msg = MessageModel(
                session_id=session_id,
                round=round_num,
                message_role=role,
                content=content,
            )
            db.add(msg)
            db.commit()
        finally:
            db.close()

    def get_next_round(self, session_id: str):
        db = self.get_session()
        try:
            last_msg = (
                db.query(MessageModel)
                .filter_by(session_id=session_id)
                .order_by(MessageModel.round.desc())
                .first()
            )
            return (last_msg.round + 1) if last_msg else 1
        finally:
            db.close()

    def update_session_title(self, session_id: str, title: str):
        db = self.get_session()
        try:
            db.query(SessionModel).filter_by(session_id=session_id).update(
                {"title": title}
            )
            db.commit()
        except Exception as e:
            db.rollback()
            raise e
        finally:
            db.close()

    def get_question_sessions(self, user_id: int, question_id: int):
        db = self.get_session()
        try:
            records = (
                db.query(SessionModel)
                .filter_by(user_id=user_id, question_id=question_id)
                .order_by(SessionModel.create_at.desc())
                .all()
            )
            return [
                {
                    "question_id": rec.question_id,
                    "session_id": rec.session_id,
                    "title": rec.title,
                }
                for rec in records
            ]
        finally:
            db.close()

    def _get_user_id_by_session_id(self, session_id: str):
        db = self.get_session()
        try:
            record = db.query(SessionModel).filter_by(session_id=session_id).first()
            return record.user_id
        finally:
            db.close()

    def get_session_message_history(self, user_id: int, session_id: str):
        db = self.get_session()
        try:
            user_id_by_session_id = self._get_user_id_by_session_id(session_id)
            if user_id_by_session_id != user_id:
                raise Exception("会话不属于当前用户")
            records = (
                db.query(MessageModel)
                .filter_by(session_id=session_id)
                .order_by(MessageModel.round)
                .offset(1)
                .all()
            )
            return [
                {
                    "role": rec.message_role,
                    "content": rec.content,
                }
                for rec in records
            ]
        finally:
            db.close()
