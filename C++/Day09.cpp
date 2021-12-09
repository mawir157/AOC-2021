#include "AH.h"

namespace Day09
{

	typedef std::vector<std::vector<int>> vector2d;

	vector2d parseInput(const std::vector<std::string>& ss)
	{
		std::vector<std::vector<int>> grid;
		for (auto s : ss)
		{
			std::vector<int> temp;
			for (auto c : s)
			{
				int v = c - 48;
				temp.push_back(v);
			}
			grid.push_back(temp);
		}

		return grid;
	}

	int part1(const vector2d& grid, std::vector<std::pair<int, int>>& ps)
	{
		int score = 0;
		for (size_t i = 0; i < grid.size(); ++i)
		{
			const auto row = grid[i];
			for (size_t j = 0; j < row.size(); ++j)
			{
				bool low = true;
				const auto value = row[j];
				if (i > 0)
					low &= (grid[i - 1][j] > value);
				if (i < 99)
					low &= (grid[i + 1][j] > value);
				if (j > 0)
					low &= (grid[i][j - 1] > value);
				if (j < 99)
					low &= (grid[i][j + 1] > value);

				if (low) {
					score += (value + 1);
					ps.push_back(std::make_pair(i, j));
				}
			}
		}

		return score;
	}

	int floodFill(vector2d& grid, const std::pair<int, int> p)
	{
		if ((p.first < 0) || (p.first > 99) || (p.second < 0) || (p.second > 99))
			return 0;

		if (grid[p.first][p.second] == 9)
			return 0;

		int size = 1;
		grid[p.first][p.second] = 9;
		size += floodFill(grid, std::make_pair(p.first - 1, p.second));
		size += floodFill(grid, std::make_pair(p.first + 1, p.second));
		size += floodFill(grid, std::make_pair(p.first, p.second - 1));
		size += floodFill(grid, std::make_pair(p.first, p.second + 1));

		return size;
	}

	int part2(vector2d& grid, const std::vector<std::pair<int, int>>& ps)
	{
		std::vector<int> basins;
		for (const auto p : ps)
			basins.push_back(floodFill(grid, p));

		std::sort(basins.begin(), basins.end(), std::greater<int>());

		return basins[0] * basins[1] * basins[2];
	}

	int Run(const std::string& filename)
	{
		const auto inputLines = AH::ReadTextFile(filename);
		auto oceanFloor = parseInput(inputLines);

		std::vector<std::pair<int, int>> lows;
		const auto p1 = part1(oceanFloor, lows);
		
		AH::PrintSoln(9, p1, part2(oceanFloor, lows));

		return 0;
	}

}
