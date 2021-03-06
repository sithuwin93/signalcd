import 'dart:async';
import 'dart:convert';
import 'dart:html';

import 'package:angular/angular.dart';
import 'package:ui/deployments_service.dart';
import 'package:ui/duration_component.dart';
import 'package:ui/src/api/api.dart';
import 'package:ui/status_component.dart';
import 'package:ui/timeago_component.dart';

@Component(
  selector: 'deployments-list',
  templateUrl: 'deployments_component.html',
  directives: [
    coreDirectives,
    StatusComponent,
    TimeagoComponent,
    DurationComponent,
  ],
  providers: [
    DeploymentsService,
  ],
)
class DeploymentsComponent implements OnInit, OnDestroy {
  final DeploymentsService _deploymentsService;

  DeploymentsComponent(this._deploymentsService);

  Timer timer;
  EventSource events = EventSource("/api/v1/deployments/events");
  List<SignalcdDeployment> deployments = [];

  @override
  void ngOnInit() {
    onDeploymentEvent(events.onMessage);
    getDeployments();

    timer = Timer.periodic(
      Duration(seconds: 30),
      (Timer timer) => getDeployments(),
    );
  }

  @override
  void ngOnDestroy() {
    timer.cancel();
    events.close();
  }

  void getDeployments() {
    _deploymentsService.deployments().then((List<SignalcdDeployment> deployments) {
      // Only update if number of deployments changed
      if (this.deployments.length != deployments.length) {
        this.deployments = deployments;
      }
    });
  }

  void onDeploymentEvent(Stream<MessageEvent> events) {
    events.forEach((MessageEvent message) {
      SignalcdDeployment deployment = SignalcdDeployment.fromJson(json.decode(message.data));

      int index = -1;

      int i = 0;
      deployments.forEach((SignalcdDeployment d) {
        if (deployment.number == d.number) {
          index = i;
        }
        i++;
      });

      if (index == -1) {
        deployments.insert(0, deployment);
      } else {
        deployments.replaceRange(index, index+1, [deployment]);
      }
    });
  }
}
