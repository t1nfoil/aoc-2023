import math
import argparse

CLI=argparse.ArgumentParser()
CLI.add_argument(
  "--rcmlist",
  nargs="*",
  type=int,
)

args= CLI.parse_args()
lcm_steps = math.lcm(*args.rcmlist)

print(lcm_steps)
