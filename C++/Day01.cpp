#include "AH.h"

namespace Day01
{

	int diff_off(const std::vector<int>& v, const size_t offset)
	{
		int acc = 0;
		for (size_t i = 0; i < v.size() - offset; ++i)
			if (v[i] < v[i + offset])
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

		AH::PrintSoln(1, diff_off(values, 1), diff_off(values, 3));

		return 0;
	}

}
