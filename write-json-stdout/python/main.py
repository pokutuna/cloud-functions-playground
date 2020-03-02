import os
import json
from flask import jsonify

def App(request):
    payload = {
        "runtime": os.getenv("GCF_RUNTIME", ""),
        "key": "value",
        "array": [1, 2, 3],
    }
    print(json.dumps(payload))
    return jsonify(payload)
