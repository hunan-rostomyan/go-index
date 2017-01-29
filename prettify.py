#!/usr/bin/env python
import json

index = json.load(open('index.json'))
with open('index.json', 'w') as f:
	json.dump(index, f, sort_keys=True, indent=2)
