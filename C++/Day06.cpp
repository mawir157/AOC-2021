#include "AH.h"

namespace Day06
{

	std::map<int, int64_t> parseInput (const std::string& s)
	{
		const auto parts = AH::Split(s, ',');
		std::map<int, int64_t> fs;
		for (const auto s : parts)
		{
			const int temp = std::stoi(s);
			fs[temp] += 1;
		}

		return fs;
	}

	std::map<int, int64_t> oneDay(const std::map<int, int64_t>& fs)
	{
		std::map<int, int64_t> gs;
		for (auto const& [k, v] : fs)
		{
			if (k == 0)
			{
				gs[6] += v;
				gs[8] += v;
			}
			else
			{
				gs[k-1] += v;
			}
		}

		return gs;
	}

	int64_t life(const std::map<int, int64_t>& fs, const int n)
	{
		auto gs = fs;
		for (int i = 0; i < n; ++i)
			gs = oneDay(gs);

		int64_t t = 0;
		for (auto const& [k, v] : gs)
			t += v;

		return t;
	}

	int Run(const std::string& filename)
	{
		auto inputLines = AH::ReadTextFile(filename);
		const auto fs = parseInput(inputLines[0]);

		AH::PrintSoln(6, life(fs, 80), life(fs, 256));

		return 0;
	}

}
