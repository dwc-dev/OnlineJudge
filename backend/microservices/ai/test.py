from ai.ai import ChatClient

api_key = "sk-a3105d444d4d49a8b2c551465ebb9ff1"
base_url = "https://api.deepseek.com"
model = "deepseek-chat"

chat_client = ChatClient(api_key, base_url, model)

print(chat_client.generate_chat_title("C++实现这道题"))
