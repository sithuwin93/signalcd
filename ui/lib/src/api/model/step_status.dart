part of swagger.api;

class StepStatus {
  
  DateTime started = null;
  

  DateTime finished = null;
  

  String phase = null;
  //enum phaseEnum {  unknown,  success,  failure,  progress,  pending,  killed,  };

  String logs = null;
  
  StepStatus();

  @override
  String toString() {
    return 'StepStatus[started=$started, finished=$finished, phase=$phase, logs=$logs, ]';
  }

  StepStatus.fromJson(Map<String, dynamic> json) {
    if (json == null) return;
    started = json['started'] == null ? null : DateTime.parse(json['started']);
    finished = json['finished'] == null ? null : DateTime.parse(json['finished']);
    phase =
        json['phase']
    ;
    logs =
        json['logs']
    ;
  }

  Map<String, dynamic> toJson() {
    return {
      'started': started == null ? '' : started.toUtc().toIso8601String(),
      'finished': finished == null ? '' : finished.toUtc().toIso8601String(),
      'phase': phase,
      'logs': logs
     };
  }

  static List<StepStatus> listFromJson(List<dynamic> json) {
    return json == null ? new List<StepStatus>() : json.map((value) => new StepStatus.fromJson(value)).toList();
  }

  static Map<String, StepStatus> mapFromJson(Map<String, Map<String, dynamic>> json) {
    var map = new Map<String, StepStatus>();
    if (json != null && json.length > 0) {
      json.forEach((String key, Map<String, dynamic> value) => map[key] = new StepStatus.fromJson(value));
    }
    return map;
  }
}

