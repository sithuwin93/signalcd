import 'dart:async';

import 'package:angular/angular.dart';
import 'package:ui/agents_service.dart';
import 'package:ui/src/api/api.dart';

@Component(
  selector: 'agents-list',
  templateUrl: 'agents_component.html',
  directives: [coreDirectives],
  providers: [AgentsService],
  pipes: [],
)
class AgentsComponent implements OnInit, OnDestroy {
  AgentsComponent(this._agentsService);

  final AgentsService _agentsService;
  Timer timer;

  List<Agent> agents = [];

  @override
  void ngOnInit() {
    getAgents();

    timer = Timer.periodic(
      Duration(seconds: 10),
      (Timer timer) => getAgents(),
    );
  }

  @override
  void ngOnDestroy() {
    timer.cancel();
  }

  void getAgents() {
    _agentsService.agents().then((List<Agent> agents) {
      this.agents = agents;
    });
  }
}
