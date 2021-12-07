#include "AH.h"

namespace Day07
{

	std::pair<int, int> maxMin(const std::vector<int>& vec)
	{
		int maxPair = -(1 << 30);
		int minPair = (1 << 30);	

		for (auto v : vec)
		{
			if (v > maxPair)
				maxPair = v;

			if (v < minPair)
				minPair = v;
		}	
		return std::make_pair(maxPair, minPair);
	}

	std::pair<int, int> crabEnergy(const std::vector<int>& crabs)
	{
		int energy1 = (1 << 30);
		int energy2 = (1 << 30);

		auto [hi, lo] = maxMin(crabs);

		for (int i = lo; i <= hi; ++i)
		{
			int temp1 = 0;
			int temp2 = 0;
			for (auto c : crabs)
			{
				int n = std::abs(c - i);
				temp1 += n;
				temp2 += (n * (n + 1)) / 2;
			}

			if (temp1 < energy1)
				energy1 = temp1;

			if (temp2 < energy2)
				energy2 = temp2;
		}

		return std::make_pair(energy1, energy2);
	}

	int Run(const std::string& filename)
	{
		auto inputLines = AH::ReadTextFile(filename);
		auto ss = AH::Split(inputLines[0], ',');
		
		std::vector<int> crabs;
		std::transform(ss.begin(), ss.end(),
		               std::back_inserter(crabs),
		               [](std::string s) -> int { return std::stoi(s); });

		auto [p1, p2] = crabEnergy(crabs);

		AH::PrintSoln(7, p1, p2);

		return 0;
	}

}
