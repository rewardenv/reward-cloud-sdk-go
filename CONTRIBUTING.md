# Contributing

```
# Copy Reward Cloud's API.json to the folder

# Install dependencies
pip3 install -r hack/requirements.txt

# Convert api.json to yaml
make convert

# Remove jsonld entries
make remove-jsonld

# Generate Go code
make generate
```