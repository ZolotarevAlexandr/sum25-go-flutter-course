import 'package:flutter/material.dart';

class CounterApp extends StatefulWidget {
  const CounterApp({super.key});

  @override
  State<CounterApp> createState() => _CounterAppState();
}

class _CounterAppState extends State<CounterApp> {
  int _counter = 0;

  void _increment() {
    setState(() {
      _counter++;
    });
  }

  void _decrement() {
    setState(() {
      _counter--;
    });
  }

  void _reset() {
    setState(() {
      _counter = 0;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      alignment: Alignment.center,
      color: Colors.white,
      child: Column(
        children: [
          const Text("Counter: "),
          Text("$_counter"),
          Row(
            children: [
              ElevatedButton(
                  onPressed: _increment, child: const Icon(Icons.add)),
              ElevatedButton(
                  onPressed: _decrement, child: const Icon(Icons.remove)),
              ElevatedButton(
                  onPressed: _reset, child: const Icon(Icons.refresh)),
              ElevatedButton(onPressed: () {}, child: const Text("Counter"))
            ],
          )
        ],
      ),
    );
  }
}
