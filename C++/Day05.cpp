#include "AH.h"

namespace Day05
{

	typedef std::pair<int, int> Point;
	typedef	std::pair<Point, Point> Line ;

	Line parseLine(const std::string& s)
	{
		auto parts = AH::SplitOnString(s, " -> ");
		auto lhs = AH::SplitOnString(parts[0], ",");
		auto rhs = AH::SplitOnString(parts[1], ",");

		Point from(std::stoi(lhs[0]), std::stoi(lhs[1]));
		Point to(std::stoi(rhs[0]), std::stoi(rhs[1]));

		Line l(from, to);

		return l;
	}

	void addLine(std::map<Point, int>& m, const Line& l, const bool d)
	{
		auto v = l.first;
		if (l.first.first == l.second.first) // second varies
		{
			const auto step = (l.second.second > l.first.second) ? 1 : -1;
			for (; v.second != l.second.second; v.second += step)
				m[v]++;
		}
		else if (l.first.second == l.second.second) // first varies
		{
			const auto step = (l.second.first > l.first.first) ? 1 : -1;
			for (; v.first != l.second.first; v.first += step)
				m[v]++;		
		}
		else if (d) // this is a diagonal pipe;
		{
			const auto step0 = (l.second.first > l.first.first) ? 1 : -1;
			const auto step1 = (l.second.second > l.first.second) ? 1 : -1;
			for (; (v.first != l.second.first) || (v.second != l.second.second); )
			{
				m[v]++;
				v.first += step0;
				v.second += step1;
			}
		}
		m[v]++;

		return;
	}

	std::pair<int, int> addLines(const std::vector<Line> ls)
	{
		std::map<Point, int> m1;
		std::map<Point, int> m2;

		std::pair<int, int> soln(0,0);

		for (const auto l : ls)
		{
			addLine(m1, l, false);
			addLine(m2, l, true);
		}

		for (auto const& x : m1)
			if (x.second > 1)
				soln.first += 1;

		for (auto const& x : m2)
			if (x.second > 1)
				soln.second += 1;

		return soln;
	}

	int Run(const std::string& filename)
	{
		auto inputLines = AH::ReadTextFile(filename);
		
		// convert lines to int
		std::vector<Line> lines;
		std::transform(inputLines.begin(), inputLines.end(),
		               std::back_inserter(lines),
		               [](std::string s) -> Line { return parseLine(s); });

		auto [p1, p2] = addLines(lines);

		AH::PrintSoln(5, p1, p2);

		return 0;
	}

}
