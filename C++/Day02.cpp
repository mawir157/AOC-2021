#include "AH.h"

namespace Day02
{

	int Run(const std::string& filename)
	{
		auto inputLines = AH::ReadTextFile(filename);

		int64_t p1v = 0, p1h = 0, p2v = 0;

		for (auto s : inputLines) {
			const auto ss = AH::Split(s, ' ');
			const int val = std::stoi(ss[1]);

			if (ss[0] == "forward") {
				p1h += val;
				p2v += p1v * val;
			} else if (ss[0] == "up" ) {
				p1v -= val;
			} else if (ss[0] == "down") {
				p1v += val;
			}
		}

		AH::PrintSoln(2, p1h * p1v, p1h * p2v);

		return 0;
	}

}
