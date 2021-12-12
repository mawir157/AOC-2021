import AdventHelper

import qualified Data.Map as Map
import Data.Char
import Data.List
import Data.List.Split
import Data.Maybe

type Edge = (String, String)

type Path = (Map.Map String Int, Bool)

parseInput :: String -> Edge
parseInput s = (ps!!0, ps!!1)
  where ps = splitOn "-" s

initPath :: [Edge] -> Bool -> Path
initPath es b = (Map.insert "start" 2 (Map.fromList xs), b)
  where cv = nub ((map fst es) ++ (map snd es))
        xs = zip cv $ repeat 0

adj :: [Edge] -> Path -> String -> [String]
adj g (caves, bonus) vert = nbr
  where allowed = if' bonus 1 0
        edf = map snd $ filter (\(f, _) -> (f == vert)) g
        edt = map fst $ filter (\(_, t) -> (t == vert)) g
        nbr = filter (\v -> (fromJust (Map.lookup v caves)) <= allowed) (edf ++ edt)

visit :: Path -> String -> Path
visit (m, b) s
  | not lc    = (m, b)
  | otherwise = (m', b')
  where lc = s == (map toLower s)
        m' = Map.adjustWithKey (\_ x -> x + 1) s m
        b' = if' (lc && ((fromJust (Map.lookup s m')) > 1)) False b

countPaths :: [Edge] -> (Path, (String, String)) -> Int
countPaths g (path, (from, to))
  | from == to = 1
  | otherwise  = sum $ map (countPaths g) $ zip paths ends
  where nbrs = adj g path from
        paths = map (visit path) nbrs
        ends = [ (v, to) | v <- nbrs ]

main = do
  putStrLn "Day 12"
  f <- readFile "../input/input12.txt"
  let es = map parseInput $ lines f

  printSoln 1 $ countPaths es ((initPath es False), ("start", "end"))
  printSoln 2 $ countPaths es ((initPath es True), ("start", "end"))
