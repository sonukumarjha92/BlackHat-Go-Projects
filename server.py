import asyncio
import json
import os
import websockets

connected_users = {}


async def broadcast_user_list():
  user_list = list(connected_users.keys())
  payload = json.dumps({"type": "user_list", "users": user_list})
  for ws in connected_users.values():
    try:
      await ws.send(payload)
    except:
      pass


async def chat_handler(websocket):
  username = None
  try:
    async for message in websocket:
      data = json.loads(message)

      if "register" in data:
        username = data["register"]
        connected_users[username] = websocket
        print(f"[+] {username} online ho gaya.")
        await broadcast_user_list()
        continue

      receiver = data.get("receiver")
      msg = data.get("message")

      if receiver in connected_users:
        outgoing = json.dumps(
            {"type": "chat", "sender": username, "message": msg}
        )
        await connected_users[receiver].send(outgoing)
      else:
        err = json.dumps({
            "type": "chat",
            "sender": "System",
            "message": f"User '{receiver}' offline hai.",
        })
        await websocket.send(err)

  except websockets.exceptions.ConnectionClosed:
    pass
  finally:
    if username and username in connected_users:
      del connected_users[username]
      print(f"[-] {username} offline ho gaya.")
      await broadcast_user_list()


async def main():
  # Cloud platform ke liye dynamic port (Render ke mutabiq)
  port = int(os.environ.get("PORT", 8000))
  server = await websockets.serve(chat_handler, "0.0.0.0", port)
  print(f"🚀 Cloud WebSocket Chat Server Port {port} par live hai...")
  await asyncio.Future()


if __name__ == "__main__":
  asyncio.run(main())

