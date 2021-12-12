#include "AH.h"

namespace Day12
{

	class Edge
	{
	public:
		std::string from;
		std::string to;

		Edge(const std::string from, const std::string to);
	};

	Edge::Edge(const std::string from, const std::string to) :
		from(from), to(to) {};

	class Graph
	{
	public:
		std::map<std::string, bool> m_vertices;
		std::vector<Edge> m_edges;
		bool m_bonus_cave;
		void addEdge(const Edge e);
		std::vector<std::string> adj(const std::string) const;
		void visit(const std::string v);
		bool done() const;

		Graph();
	private:
		std::map<std::string, int> visited;		
	};

	Graph::Graph() {}
	
	void Graph::addEdge(const Edge e)
	{
		visited[e.from] = 0;
		visited[e.to] = 0;

		auto eLower = e.from;
		for(auto& c : eLower)
			c = tolower(c);

		if (e.from == eLower)
			m_vertices[e.from] = false;
		else
			m_vertices[e.from] = true;

		eLower = e.to;
		for(auto& c : eLower)
			c = tolower(c);

		if (e.to == eLower)
			m_vertices[e.to] = false;
		else
			m_vertices[e.to] = true;

		m_edges.push_back(e);

		return;
	}

	void Graph::visit(const std::string v)
	{
		if (!m_vertices[v]) // v is lower case
		{
			visited[v]++; // flag as visited
			if (visited[v] > 1) // we are visiting a small cave twice..
				m_bonus_cave = false; // ...don' do that again!
		}
	}

	std::vector<std::string> Graph::adj(const std::string v) const
	{
		std::vector<std::string> vs;
		// how many pervious visted to a small cave are allowed?
		const int allowed = m_bonus_cave ? 1 : 0;

		for (auto e : m_edges)
		{
			if (e.from == v)
			{
				if (visited.at(e.to) <= allowed)
					vs.push_back(e.to);
			}
			else if (e.to == v)
			{
				if (visited.at(e.from) <= allowed)
					vs.push_back(e.from);
			}
		}			

		return vs;
	}

	Graph buildGraph(const std::vector<std::string>& ss, const bool part2 = false)
	{
		Graph g;
		for (auto s : ss)
		{
			auto parts = AH::Split(s, '-');
			Edge e(parts[0], parts[1]);
			g.addEdge(e);
		}
		g.visit("start");
		g.visit("start");
		g.m_bonus_cave = part2;

		return g;
	}

	int countPaths(const Graph g, const std::string from, const std::string to)
	{
		if (from == to)
			return 1;

		int routes = 0;
		auto ns = g.adj(from); // get all the neighbouring permitted vertices
		for (auto n : ns)
		{
			auto g_new = g;
			g_new.visit(n);
			routes += countPaths(g_new, n, to);
		}

		return routes;
	}

	int Run(const std::string& filename)
	{
		auto inputLines = AH::ReadTextFile(filename);
		const auto g1 = buildGraph(inputLines);
		const auto g2 = buildGraph(inputLines, true);

		AH::PrintSoln(12, countPaths(g1, "start", "end"),
			                countPaths(g2, "start", "end"));

		return 0;
	}

}
