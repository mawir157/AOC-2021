#include "AH.h"

namespace Day11
{

	struct Octopus
	{
	public:
		int Level;
		bool Flashed;

		Octopus(const int Level, const bool Flashed);
		Octopus();
	};

	Octopus::Octopus(const int Level, const bool Flashed) :
		Level(Level), Flashed(Flashed) {}
	Octopus::Octopus() : Level(-1), Flashed(false) {}

	typedef std::map<std::pair<int, int>, Octopus> Octopi;

	Octopi GetOctopi(const std::vector<std::string>& ss)
	{
		Octopi ocks;

		int i = 0, j = 0;
		for (auto s : ss)
		{
			j = 0;
			for (auto c : s)
			{
				std::pair<int, int> pos = std::make_pair(i,j);
				int value = c - 48;
				ocks[pos] = Octopus(value, false);
				j++;
			}
			i++;
		}

		return ocks;
	}

	std::vector<std::pair<int, int>> nbrs(const std::pair<int, int> p)
	{
		std::vector<std::pair<int, int>> ps;
		ps.emplace_back(p.first - 1, p.second - 1);
		ps.emplace_back(p.first - 1, p.second);
		ps.emplace_back(p.first - 1, p.second + 1);
		ps.emplace_back(p.first,     p.second - 1);
		ps.emplace_back(p.first,     p.second + 1);
		ps.emplace_back(p.first + 1, p.second - 1);
		ps.emplace_back(p.first + 1, p.second);
		ps.emplace_back(p.first + 1, p.second + 1);

		return ps;
	}

	unsigned int octoTick(Octopi& ocks)
	{
		// increment all the octopode
		for (auto & [p, o] : ocks)
				o.Level++;

		// flash all octopode
		unsigned int total_flashes = 0;
		unsigned int flash_count = 1;
		while (flash_count != 0)
		{
			flash_count = 0;
			// check if an octopus will flash
			for (auto & [p, o] : ocks)
			{
				if ((o.Level > 9) && (!o.Flashed))
				{
					++flash_count;
					o.Flashed = true;

					auto ns = nbrs(p); // get all neighbours
					for (auto n : ns)
					{
						auto nOck = ocks.find(n);
						if (nOck != ocks.end()) // if neighbour exists...
							(nOck->second).Level += 1; // ...increment it
					}
				}
			}

			total_flashes += flash_count;
		}

		// the octopi have finished flashing
		for (auto & [p, o] : ocks)
		{
			if (o.Level > 9)
				o.Level = 0;

			o.Flashed = false;
		}

		return total_flashes;
	}

	int Run(const std::string& filename)
	{
		auto inputLines = AH::ReadTextFile(filename);
		auto ocks = GetOctopi(inputLines);

		int part1 = 0;
		int part2 = 0;

		for (part2 = 0; part2 < 100; ++part2)
			part1 += octoTick(ocks);

		while (true)
		{
			++part2;
			if (octoTick(ocks) == ocks.size())
				break;
		}

		AH::PrintSoln(11, part1, part2);

		return 0;
	}

}
