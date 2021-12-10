import AdventHelper

import Data.List

score :: String -> Int
score [] = 0
score (s:ss)
  | s == ')' = 3 + score ss
  | s == ']' = 57 + score ss
  | s == '}' = 1197 + score ss
  | s == '>' = 25137 + score ss
  | s == 'X' = 0 + score ss 

score2 :: String -> Int
score2 [] = 0
score2 (s:ss) 
  | s == '(' = s' + 1
  | s == '[' = s' + 2
  | s == '{' = s' + 3
  | s == '<' = s' + 4
  where s' = 5 * score2 ss

closing :: Char -> Char
closing c
  | c == ')'  = '('
  | c == ']'  = '['
  | c == '}'  = '{'
  | c == '>'  = '<'
  | otherwise = c

firstIllegal :: (String, String) -> (Char, String)
firstIllegal ([],stack) = ('X', stack) -- no errors but incomplete
firstIllegal ((s:ss), []) = firstIllegal (ss, [s])
firstIllegal ((s:ss), (t:tt))
  | s `elem` "{([<"  = firstIllegal (ss, (s:t:tt)) -- add to stack 
  | t == (closing s) = firstIllegal (ss, tt) -- cancel out
  | otherwise        = (s, "") -- return string that caused problems

main = do
  putStrLn "Day 10"
  f <- readFile "../input/input10.txt"
  let ss = map (\s -> (s, "")) $ lines f
  let t = map (firstIllegal) ss

  let tt = filter (\s -> fst (firstIllegal s) == 'X') ss
  let x = map (score2. reverse. snd .firstIllegal) tt

  printSoln 1 $ score $ map fst t 
  printSoln 2 $ head $ drop ((length x) `div` 2) $ sort x
