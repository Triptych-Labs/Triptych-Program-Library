from datetime import datetime, timezone

timestamp = (
    datetime.now(timezone.utc)
    .replace(hour=0, minute=0, second=0, microsecond=0)
    .timestamp()
)

print(timestamp)
print(datetime.utcfromtimestamp(timestamp).isoformat())

