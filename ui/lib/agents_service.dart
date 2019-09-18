import 'package:ui/api.dart';
import 'package:ui/src/api/api.dart';

class AgentsService {
  final API _api;

  AgentsService(this._api);

  Future<List<Agent>> agents() async {
    return _api.agents.agents();
  }
}
