#include "AH.h"

namespace Day03
{

	std::vector<unsigned int> bitCount(const std::set<std::string>& s,
	                                   const size_t size)
	{
		std::vector<unsigned int> count(size);

		for (auto i = s.begin(); i != s.end(); ++i)
		{
			auto str = *i; 
			for (int i = 0; i < str.size(); i++)
				if (str[i] == '1')
					count[i] += 1;
		}

		return count;
	}

	int part1(const std::set<std::string>& s, const size_t size)
	{
		const auto mostCommon = bitCount(s, size);

		int gamma = 0, epsilon = 0;
		for (auto b : mostCommon)
		{
			gamma <<= 1;
			epsilon <<= 1;

			if (2 * b >= s.size()) 
				gamma += 1;
			else
				epsilon += 1;
		}

		return gamma * epsilon;
	}

	int stringToBin(const std::string& s)
	{
		int bin = 0;
		for (auto c : s)
		{
			bin <<= 1;
			bin += (c == '1') ? 1 : 0;
		}

		return bin;
	}

	int setReduce(const std::set<std::string>& s, const size_t size,
	              const bool good)
	{
		auto t = s;

		for (size_t i = 0; i < size; i++)
		{
			const auto mostCommon = bitCount(t, size);		

			const char r = (2 * mostCommon[i] >= t.size()) ? '1' : '0';

			for (auto it = t.begin(); it != t.end(); )
			{
				if ((it->at(i) != r) == good)
					it = t.erase(it);
				else
					++it;
			}

			if (t.size() == 1)
				return stringToBin(*(t.begin()));
		}

		return -1;
	}

	int part2(const std::set<std::string>& s, const size_t size)
	{
		return setReduce(s, size, true) * setReduce(s, size, false);
	}

	int Run(const std::string& filename)
	{
		auto inputLines = AH::ReadTextFile(filename);
		auto size = inputLines[0].size();
		std::set<std::string> s(inputLines.begin(), inputLines.end());

		AH::PrintSoln(3, part1(s, size), part2(s, size));

		return 0;
	}

}
