#include "AH.h"

namespace Day01
{

	int Part1(const std::vector<int> v)
	{
		int acc = 0;
		for (size_t i = 1; i < v.size(); ++i)
			if (v[i - 1] < v[i])
				acc += 1;

	  return acc;
	}

	int Part2(const std::vector<int> v)
	{
		int acc = 0;
		for (size_t i = 0; i < v.size() - 3; ++i)
			if (v[i] < v[i + 3])
				acc += 1;

	  return acc;
	}

	int Run(const std::string& filename)
	{
		auto inputLines = AH::ReadTextFile(filename);
		
		// convert lines to int
		std::vector<int> values;
		std::transform(inputLines.begin(), inputLines.end(),
		               std::back_inserter(values),
		               [](std::string s) -> int { return std::stoi(s); });

		AH::PrintSoln(1, Day01::Part1(values), Day01::Part2(values));

		return 0;
	}

}
