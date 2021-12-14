#include "AH.h"

namespace Day13
{

	typedef std::pair<int, int> Point;
	typedef std::pair<bool, int> Fold;

	std::set<Point> paperFold(const std::set<Point> ps, const Fold f)
	{
		std::set<Point> newPs;
		for (auto p : ps) 
		{
			Point pNew;
			if (f.first)
				pNew = std::make_pair(p.first, (p.second > f.second) ? 2*f.second - p.second : p.second);
			else
				pNew = std::make_pair((p.first > f.second) ? 2*f.second - p.first : p.first, p.second);

			newPs.insert(pNew);
		}

		return newPs;
	}

	void PrintGrid(const std::set<Point> ps)
	{
		for (int i = 0; i < 6; ++i)
		{
			for (int j = 0; j < 40; ++j)
			{
				auto p = std::make_pair(j,i);
				if (ps.count(p) != 0)
					std::cout << "#";
				else
					std::cout << " ";
			}
			std::cout << std::endl;
		}
	}

	int Run(const std::string& filename)
	{
		auto inputLines = AH::ReadTextFile(filename);

		std::set<Point> points;
		std::vector<Fold> folds;

		size_t l = 0;
		while (inputLines[l].size() > 0)
		{
			const auto ps = AH::Split(inputLines[l], ',');
			const auto p = std::make_pair(std::stoi(ps[0]),std::stoi(ps[1]));
			points.insert(p);
			++l;
		}
		++l;
		for (; l < inputLines.size(); ++l)
		{
			auto str = inputLines[l];
			const bool b = (str.at(11) == 'y');
			str.erase(0,13);
			const auto f = std::make_pair(b, std::stoi(str));
			folds.push_back(f);
		}
		auto part1 = paperFold(points, folds.front());
		std::set<Point> part2 = points;
		for (auto f : folds)
			part2 = paperFold(part2, f);

		AH::PrintSoln(13, part1.size(), part2.size());
		PrintGrid(part2);
		return 0;
	}

}
