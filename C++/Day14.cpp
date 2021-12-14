#include "AH.h"

namespace Day14
{

	typedef std::map<std::string, std::string> Rules;
	typedef std::map<std::string, uint64_t> PairCounts;

	Rules parseInput(const std::vector<std::string> ss)
	{
		std::map<std::string, std::string> m;
		for (size_t i = 2; i < ss.size(); ++i)
		{
			auto ps = AH::SplitOnString(ss[i], " -> ");
				m[ps[0]] = ps[1];
		}

		return m;
	}

	PairCounts stringToPairs(const std::string s)
	{
		PairCounts pc;
		for (size_t i = 0; i < s.size() - 1; ++i)
			pc[s.substr(i, 2)] += 1;

		return pc;
	}

	PairCounts applyRules(const Rules rs, const PairCounts ps)
	{
		PairCounts newPairs;
		for (auto const &[pair, count] : ps)
		{
			const auto r = rs.at(pair);

			newPairs[(pair.at(0) + r)] += count;
			newPairs[(r + pair.at(1))] += count;
		}

		return newPairs;
	}

	PairCounts repAapplyRules(const unsigned int n, const Rules rs,
		                        const PairCounts ps)
	{
		auto psTemp = ps;
		for (unsigned int i = 0; i < n; ++i)
			psTemp = applyRules(rs, psTemp);

		return psTemp;
	}

	int64_t score(const PairCounts ps, const char final)
	{
		uint64_t max = 0;
		uint64_t min = 1; (min <<= 62);

		std::map<char, uint64_t> charCount;
		for (auto const & [pair, count] : ps)
			charCount[pair.at(0)] += count;

		charCount[final] += 1;

		for (auto const & [pair, count] : charCount)
		{
			if (count > max)
				max = count;
			if (count < min)
				min = count;
		}

		return max - min;
	}

	int Run(const std::string& filename)
	{
		auto inputLines = AH::ReadTextFile(filename);
		auto pairs = stringToPairs(inputLines[0]);
		auto rules = parseInput(inputLines);
		
		const auto final = inputLines[0].back();
		const auto p1 = repAapplyRules(10, rules, pairs);
		const auto p2 = repAapplyRules(40, rules, pairs);

		AH::PrintSoln(14, score(p1, final), score(p2, final));

		return 0;
	}

}
