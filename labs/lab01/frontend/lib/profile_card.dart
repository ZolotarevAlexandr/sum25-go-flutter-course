import 'package:flutter/material.dart';

class ProfileCard extends StatelessWidget {
  final String name;
  final String email;
  final int age;
  final String? avatarUrl;

  const ProfileCard({
    super.key,
    required this.name,
    required this.email,
    required this.age,
    this.avatarUrl,
  });

  @override
  Widget build(BuildContext context) {
    final initial = name.isNotEmpty ? name[0].toUpperCase() : '?';

    return Padding(
        padding: const EdgeInsets.all(16),
        child: Row(
          children: [
            CircleAvatar(
              radius: 30,
              child: Text(initial, style: const TextStyle(fontSize: 24)),
            ),
            const SizedBox(width: 24),
            Column(
              children: [
                Text(name, style: const TextStyle(fontSize: 24)),
                Text(email),
                Text("Age: $age"),
              ],
            )
          ],
        ));
  }
}
