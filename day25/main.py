from networkx import Graph, minimum_edge_cut, connected_components

# Went back to python to use networkx, will return someday to find a golang sol
# I wanted to create a graph in golang originally and find 3 edges to make 2
# connected components, but didn't have the time this christmas day
# Alas, 23.5/25 days in pure golang will have to do
graph = Graph()
with open('input.txt', 'r') as file:
    for line in file:
        line = line.strip()
        wireConnections = line.split(":")[1].strip().split(" ")
        for edge in wireConnections:
            graph.add_edge(line.split(":")[0], edge)

min_edges = minimum_edge_cut(graph)
for edge in min_edges:
    graph.remove_edge(edge[0], edge[1])

components = list(connected_components(graph))

print(len(components[0])*len(components[1]))