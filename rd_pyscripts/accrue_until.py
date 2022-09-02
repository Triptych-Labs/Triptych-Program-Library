import time

###
# This script attests the following behaviour:
"""
    Assign an end time
    Sample yields over duration*
    
    * sampling after end time always yields zero
      if number of samples is greater than zero.
"""


class Quest:
    start: int
    end: int
    last: int

    def __init__(self, now, duration):
        self.start = now
        self.end = now + duration
        self.last = now

    def make_claim(self, now: int) -> int:
        if now > self.end and self.last == self.end:
            return 0

        duration = 0
        if now >= self.end:
            duration = self.end - self.last
            self.last = self.end
        else:
            duration = now - self.last
            self.last = now

        return 5 * duration


if __name__ == "__main__":
    now = int(time.time())

    quest = Quest(now, 10)

    acc = 0
    for _ in range(5):
        time.sleep(3)
        claim = quest.make_claim(int(time.time()))
        acc += claim
        print(claim, acc)

    print("finished", acc)

