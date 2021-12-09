#include "AH.h"

namespace Day08
{

	void parseLine(const std::string& s, std::vector<std::string>& messed,
		             std::vector<std::string>& codes)
	{
		auto ps = AH::SplitOnString(s, " | ");
		messed = AH::Split(ps[0], ' ');
		codes = AH::Split(ps[1], ' ');

		for (auto& s : messed)
			std::sort(s.begin(), s.end());

		for (auto& s : codes)
			std::sort(s.begin(), s.end());

		return;
	}

	int part1(const std::vector<std::string>& ss)
	{
		int total = 0;
		for (auto s : ss)
			if ( (s.size() == 2) || (s.size() == 3) || (s.size() == 4) || (s.size() == 7) )
				++total;

		return total;
	}

	std::string stringDiff(const std::string& a, const std::string& b)
	{
		std::string diff = "";
		for (auto c : a)
			if (b.find(c) == std::string::npos)
				diff += c;

		return diff;
	}

	std::map<int, std::string> unscrable(const std::vector<std::string>& messed)
	{
		std::map<int, std::string> m;
		std::set<std::string> ss;

		for (auto s : messed)
			ss.insert(s);

		for (auto it = ss.begin(); it != ss.end(); )
		{
			auto s = (*it);
			if (s.size() == 2)
			{
				m[1] = s;
				it = ss.erase(it);
			}
			else if (s.size() == 3)
			{
				m[7] = s;
				it = ss.erase(it);
			}
			else if (s.size() == 4)
			{
				m[4] = s;
				it = ss.erase(it);
			}
			else if (s.size() == 7)
			{
				m[8] = s;
				it = ss.erase(it);	
			}
			else
			{
				++it;
			}
		}

		// length 6 words
		for (auto s : ss)
		{
			if (s.size() != 6)
				continue;

			if (stringDiff(m[4], s).size() == 0)
			{
				m[9] = s;
				ss.erase(s);
				break;
			}
		}

		for (auto it = ss.begin(); it != ss.end(); )
		{
			auto s = (*it);

			if (s.size() != 6)
			{
				++it;
				continue;
			}

			if (stringDiff(m[1], s).size() == 0)
				m[0] = s;	
			else
				m[6] = s;

			it = ss.erase(it);
		}

		// length 5 words
		for (auto s : ss)
		{
			if (s.size() != 5)
				continue;

			if (stringDiff(m[1], s).size() == 0)
			{
				m[3] = s;
				ss.erase(s);
				break;
			}
		}

		for (auto it = ss.begin(); it != ss.end(); )
		{
			auto s = (*it);

			if (s.size() != 5)
			{
				++it;
				continue;
			}

			if (stringDiff(s, m[9]).size() == 0)
				m[5] = s;	
			else
				m[2] = s;

			it = ss.erase(it);
		}

		return m;
	}

	int match(const std::map<int, std::string>& m, std::vector<std::string>& cs)
	{
		int decode = 0;
		for (auto code : cs)
			for (auto [key, val] : m)
				if (code == val) {
					decode *= 10;
					decode += key;
				}
		return decode;
	}

	int Run(const std::string& filename)
	{
		auto inputLines = AH::ReadTextFile(filename);

		int total1 = 0, total2 = 0;
		for (auto s : inputLines)
		{
			std::vector<std::string> messed;
			std::vector<std::string> codes;
			parseLine(s, messed, codes);
			
			total1 += part1(codes);
			
			const auto ordered = unscrable(messed);
			total2 += match(ordered, codes);
		}
		
		AH::PrintSoln(8, total1, total2);

		return 0;
	}

}
