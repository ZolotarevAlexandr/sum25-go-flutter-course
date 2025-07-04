import 'dart:async';

// ChatService handles chat logic and backend communication
class ChatService {
  final StreamController<String> _controller =
      StreamController<String>.broadcast();
  bool failConnect = false;
  bool failSend = false;

  ChatService();

  Future<void> connect() async {
    if (failConnect) {
      throw Exception("connection failed");
    }
    await Future.delayed(Duration(milliseconds: 500));
  }

  Future<void> sendMessage(String msg) async {
    if (failSend) {
      throw Exception("send failed");
    }
    await Future.delayed(Duration(milliseconds: 500));
    _controller.add(msg);
  }

  Stream<String> get messageStream {
    return _controller.stream;
  }
}
