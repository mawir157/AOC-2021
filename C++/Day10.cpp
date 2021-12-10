#include "AH.h"

namespace Day10
{

	int score1(const char c)
	{
		switch(c) {
			case ')':
				return 3;
			case ']':
				return 57;
			case '}':
				return 1197;
			case '>':
				return 25137;
			default:
				return 0;
		} 
	}

	int64_t score2(const std::string s)
	{
		int64_t score = 0;
		for (int i = s.length() - 1; i >= 0; --i)
		{
			score *= 5;
			switch(s[i]) {
				case '(':
					score += 1; break;
				case '[':
					score += 2; break;
				case '{':
					score += 3; break;
				case '<':
					score += 4; break;
				default:
					return -1;
			} 			
		}
		return score;
	}

	bool closing(const char l, const char r)
	{
		return ((l == '(' && r == ')') || (l == '{' && r == '}') ||
		        (l == '[' && r == ']') || (l == '<' && r == '>'));
	}

	std::pair<char, std::string> parseString(const std::string& s)
	{
		std::string stck = "";
		const std::string bras = "({[<";
		for (auto c : s)
		{
			if (bras.find(c) != std::string::npos)
				stck += c;
			else if (closing(stck.back(), c))
				stck.pop_back(); 
			else
				return std::make_pair(c, "");
		}

		return std::make_pair('X', stck);
	}

	int Run(const std::string& filename)
	{
		auto inputLines = AH::ReadTextFile(filename);

		int part1 = 0;
		std::vector<int64_t> part2Arr;
		int goodCount = 0;
		for (auto s : inputLines)
		{
			auto [a, b] = parseString(s);
			if (a == 'X') {
				++goodCount;
				part2Arr.push_back(score2(b));
			}

			part1 += score1(a);
		}
		std::sort(part2Arr.begin(), part2Arr.end());

		AH::PrintSoln(10, part1, part2Arr[goodCount / 2]);

		return 0;
	}

}
