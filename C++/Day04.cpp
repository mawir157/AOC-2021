#include "AH.h"

namespace Day04
{

	typedef std::vector<int> Board;

	void callNumber(Board& b, const int n)
	{
		for (size_t i = 0; i < b.size(); ++i)
			if (b[i] == n)
				b[i] = -1;
	}

	bool isWinning(const Board& b)
	{
		for (size_t i = 0; i < b.size(); i += 5)
		{
			int total = 0;
			for (size_t j = 0; j < 5; ++j)
				total += b[i + j];

			if (total == -5)
				return true;
		}

		
		for (size_t j = 0; j < 5; ++j)
		{
			int total = 0;
			for (size_t i = 0; i < b.size(); i += 5)
				total += b[i + j];

			if (total == -5)
				return true;
		}

		return false;
	}

	int score(const Board& b)
	{
		int s = 0;
		for (auto v : b)
			s += (v > 0) ? v : 0;

		return s;
	}

	std::vector<int> parseNumbers(const std::string& s)
	{
		auto strInts = AH::Split(s, ',');
		std::vector<int> values;
		std::transform(strInts.begin(), strInts.end(),
		               std::back_inserter(values),
		               [](std::string s) -> int { return std::stoi(s); });

		return values;
	}

	std::vector<Board> parseBoards(const std::vector<std::string>& ss)
	{
		int counter = 0;
		std::vector<Board> boards;
		Board temp;
		for (size_t i = 2; i < ss.size(); ++i)
		{
			auto lines = AH::Split(ss[i], ' ');
			for (auto l : lines) {
				if (l.size() == 0)
					continue;

				temp.push_back(std::stoi(l));
			}
			++counter;
			if (counter == 5) {
				boards.push_back(temp);
				temp.clear();
				counter = 0;
				++i;
			}
		}
		return boards;
	}

	int part1(std::vector<Board>& boards, const std::vector<int>& calls)
	{
		for (auto call : calls)
		{
			for (auto& board : boards)
			{
				callNumber(board, call);
				if (isWinning(board)) {
					return call * score(board);
				}
			}
		}
		return -1;
	}

	int part2(std::vector<Board>& boards, const std::vector<int>& calls)
	{
		size_t last_c = 0;
		size_t last_b = 0;

		for (size_t ib = 0; ib < boards.size(); ++ib)
		{
			for (size_t ic = 0;  ic < calls.size(); ++ic)
			{
				callNumber(boards[ib], calls[ic]);
				if (isWinning(boards[ib]))
				{
					if (ic > last_c)
					{
						last_c = ic;
						last_b = ib;
					}
					break;
				}
			}
		}

		return score(boards[last_b]) * calls[last_c];
	}

	int Run(const std::string& filename)
	{
		auto inputLines = AH::ReadTextFile(filename);
		const auto calls = parseNumbers(inputLines[0]);
		auto boards = parseBoards(inputLines);

		const auto p1 = part1(boards, calls);
		const auto p2 = part2(boards, calls);

		AH::PrintSoln(4, p1, p2);

		return 0;
	}

}
