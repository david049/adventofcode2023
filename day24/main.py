from z3 import Solver, BitVec, sat

# figured out how to use z3 just for part 2..

hailStones = []
with open('input.txt', 'r') as file:
    for line in file:
        line = line.strip().split('@')
        position = tuple(map(int, line[0].split(',')))
        velocity = tuple(map(int, line[1].split(',')))
        result = position + velocity
        hailStones.append(result)

solver = Solver()
x = BitVec('x', 64)
y = BitVec('y', 64)
z = BitVec('z', 64)
vx = BitVec('vx', 64)
vy = BitVec('vy', 64)
vz = BitVec('vz', 64)
for i, hailstone in enumerate(hailStones[:5]): # try solving for just the first 5
    (positionX, positionY, positionZ, velocityX, velocityY, velocityZ) = hailstone
    t = BitVec(f"t{i}", 64)
    solver.add(x + vx * t == positionX + velocityX * t)
    solver.add(y + vy * t == positionY + velocityY * t)
    solver.add(z + vz * t == positionZ + velocityZ * t)
if solver.check() == sat:
    model = solver.model()
    x = model.eval(x)
    y = model.eval(y)
    z = model.eval(z)
    print(x.as_long()+y.as_long()+z.as_long())
