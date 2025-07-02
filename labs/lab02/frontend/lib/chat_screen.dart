import 'package:flutter/material.dart';
import 'chat_service.dart';
import 'dart:async';

// ChatScreen displays the chat UI
class ChatScreen extends StatefulWidget {
  final ChatService chatService;

  const ChatScreen({super.key, required this.chatService});

  @override
  State<ChatScreen> createState() => _ChatScreenState();
}

class _ChatScreenState extends State<ChatScreen> {
  final TextEditingController _textController = TextEditingController();

  final List<String> _messages = [];
  bool _isLoading = false;
  String? _error;

  StreamSubscription<String>? _subscription;

  @override
  void initState() {
    super.initState();
    _initChat();
  }

  Future<void> _initChat() async {
    setState(() {
      _isLoading = true;
    });

    try {
      await widget.chatService.connect();
    } catch (e) {
      setState(() {
        _error = e.toString();
        _isLoading = false;
      });
      return;
    }

    _subscription = widget.chatService.messageStream.listen((message) {
      setState(() {
        _messages.add(message);
      });
    }, onError: (err) {
      setState(() {
        _error = err.toString();
      });
    });

    setState(() {
      _isLoading = false;
    });
  }

  @override
  void dispose() {
    _subscription?.cancel();
    _textController.dispose();
    super.dispose();
  }

  void _sendMessage() async {
    final text = _textController.text;
    if (text.isEmpty) {
      return;
    }

    try {
      await widget.chatService.sendMessage(text);
      _textController.clear();
    } catch (e) {
      _error = e.toString();
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Chat')),
      body: Column(
        children: [
          if (_isLoading) CircularProgressIndicator(),
          if (_error != null) Text("Connection error: $_error"),
          Expanded(
              child: ListView.builder(
            itemCount: _messages.length,
            itemBuilder: (context, index) {
              return ListTile(
                title: Text(_messages[index]),
              );
            },
          )),
          Padding(
            padding: EdgeInsets.all(8.0),
            child: Row(
              children: [
                Expanded(
                    child: TextField(
                  controller: _textController,
                  decoration: const InputDecoration(hintText: "Input message"),
                )),
                IconButton(
                  onPressed: _sendMessage,
                  icon: Icon(Icons.send),
                )
              ],
            ),
          )
        ],
      ),
    );
  }
}
