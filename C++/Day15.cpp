#include "AH.h"

namespace Day15
{

	typedef std::pair<uint, uint> Pos;

	struct pair_hash {
			std::size_t operator () (const std::pair<uint, uint> &p) const {
				return 10000 * p.first + p.second;
		}
	};

	std::map<Pos, uint>parseInput(const std::vector<std::string> ss)
	{
		std::map<Pos, uint>cave;

		for (size_t i = 0; i < ss.size(); ++i)
		{
			for (size_t j = 0; j < ss[i].size(); ++j)
			{
				const Pos p = std::make_pair(i, j);
				const auto v = std::stoi(ss[i].substr(j, 1));
				cave[p] = v;
			}
		}

		return cave;
	}

	std::map<Pos, uint>expandCave(const std::map<Pos, uint>cave,
	                              const uint n, const uint size)
	{
		std::map<Pos, uint>new_cave;
		for (size_t i = 0; i < n; ++i)
		{
			for (size_t j = 0; j < n; ++j)
			{
				for (auto [pos, val] : cave)
				{
					auto new_val = (val + i + j);
					new_val -= (new_val > 9) ? 9 : 0;

					Pos new_pos = std::make_pair(pos.first + (i*size),
					                             pos.second + (j*size));

					new_cave[new_pos] = new_val;
				}
			}
		}
		return new_cave;
	}

	std::vector<std::vector<bool>> initBoolGrid(const uint i)
	{
		std::vector<std::vector<bool>> grid(i, std::vector<bool>(i, true));

		return grid;
	}

	std::vector<std::vector<uint>> initIntGrid(const uint i)
	{
		std::vector<std::vector<uint>> grid(i, std::vector<uint>(i, 0));

		return grid;
	}


	Pos minDist(const std::vector<std::vector<uint>>& dist,
	            const std::unordered_set<Pos, pair_hash>& Q)
	{
		uint min = 1000000 - 1;
		Pos p;
		for (auto & q : Q)
		{
			const auto v = dist[q.first][q.second];
			if (v < min)
			{
				min = v;
				p = q;
			}
		}
		return p;
	}

	std::vector<Pos> nbrs(const uint dim, const Pos& p,
	                      const std::vector<std::vector<bool>>& Q)
	{
		std::vector<Pos> ns;
		ns.reserve(4);

		if (p.second > 0) {
			auto pu = std::make_pair(p.first, p.second - 1);
			if (Q[p.first][p.second - 1])
				ns.push_back(pu);
		}

		if (p.first > 0) {
			auto pl = std::make_pair(p.first - 1, p.second);
			if (Q[p.first - 1][p.second])
				ns.push_back(pl);
		}

		if (p.second < (dim - 1)) {
			auto pd = std::make_pair(p.first, p.second + 1);
			if (Q[p.first][p.second + 1])
				ns.push_back(pd);
		}

		if (p.first < (dim - 1)) {
			auto pu = std::make_pair(p.first + 1, p.second);
			if (Q[p.first + 1][p.second])
				ns.push_back(pu);
		}

		return ns;
	}

	uint dij(const std::map<Pos, uint>g, const Pos source, const Pos target,
	         const uint dim)
	{
		auto visited = initBoolGrid(dim);
		std::unordered_set<Pos, pair_hash> Flagged;
		auto dist = initIntGrid(dim);

		for (auto [k,v] : g)
			dist[k.first][k.second] = 1000000;

		dist[source.first][source.second] = 0;
		Flagged.insert(source);

		while (true)
		{
			auto u = minDist(dist, Flagged);

			visited[u.first][u.second] = false;
			Flagged.erase(u);
			if (u == target)
				return dist[u.first][u.second];

			auto ns = nbrs(dim, u, visited);
			for (auto & n : ns)
			{
				auto alt = dist[u.first][u.second] + g.at(n);
				if (alt < dist[n.first][n.second])
					dist[n.first][n.second] = alt;

				Flagged.insert(n);
			}
		}

		return 0;
	}

	int Run(const std::string& filename)
	{
		auto inputLines = AH::ReadTextFile(filename);
		const auto cave = parseInput(inputLines);

		const Pos start = std::make_pair(0, 0 );
		const Pos target1 = std::make_pair(99,99);
		const auto dist1 = dij(cave, start, target1, 100);

		const Pos target2 = std::make_pair(499, 499);
		const auto big_cave = expandCave(cave, 5, 100);
		const auto dist2 = dij(big_cave, start, target2, 500);

		AH::PrintSoln(15, dist1, dist2);

		return 0;
	}

}
