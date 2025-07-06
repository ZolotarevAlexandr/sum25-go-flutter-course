import 'package:flutter/material.dart';
import 'package:lab02_chat/user_service.dart';

// UserProfile displays and updates user info
class UserProfile extends StatefulWidget {
  final UserService
      userService; // Accepts a user service for fetching user info
  const UserProfile({super.key, required this.userService});

  @override
  State<UserProfile> createState() => _UserProfileState();
}

class _UserProfileState extends State<UserProfile> {
  Map<String, String>? _user;
  bool _isLoading = false;
  String? _error;

  Future<void> _fetchUser() async {
    setState(() {
      _isLoading = true;
      _error = null;
    });

    try {
      final user = await widget.userService.fetchUser();
      setState(() {
        _user = user;
        _isLoading = false;
      });
    } catch (e) {
      setState(() {
        _error = e.toString();
        _isLoading = false;
      });
    }
  }

  @override
  void initState() {
    super.initState();
    _fetchUser();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('User Profile')),
      body: Column(
        children: [
          if (_isLoading) Center(child: CircularProgressIndicator()),
          if (_error != null) Center(child: Text("error: $_error")),
          if (_user == null && !_isLoading && _error == null)
            Center(child: const Text("No user data")),
          if (_user != null)
            Center(
              child: Column(
                children: [
                  Text(_user!['name']!),
                  const SizedBox(height: 8.0),
                  Text(_user!['email']!)
                ],
              ),
            )
        ],
      ),
    );
  }
}
