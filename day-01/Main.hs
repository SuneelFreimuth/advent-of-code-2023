import Data.Char (isDigit, digitToInt)
import Data.List (find, isPrefixOf)
import Data.Maybe (fromJust)

main = do
  contents <- readFile "input.txt"
  _ <- writeFile "output.txt" ""
  let lines_ = filter (/="") $ lines contents
  let values = map callibrationValue lines_
  let formatLine (value, line) = show value ++ ": " ++ line ++ "\n"
  mapM_ (print . allDigits) lines_
  mapM_ (appendFile "output.txt" . formatLine) $ zip values lines_
  print (sum values)

callibrationValue :: String -> Int
callibrationValue line = 
  let
    digits = allDigits line
  in
    (10 * head digits) + last digits

allDigits :: String -> [Int]
allDigits [] = []
allDigits str =
  let
    (mdigit, rest) = takeDigit str
  in
    case mdigit of
      Just digit -> digit : allDigits rest
      Nothing -> allDigits rest

takeDigit :: String -> (Maybe Int, String)
takeDigit str
  | isDigit (head str) = (Just (digitToInt $ head str), tail str)
  | "one" `isPrefixOf` str = (Just 1, drop 3 str)
  | "two" `isPrefixOf` str = (Just 2, drop 3 str)
  | "three" `isPrefixOf` str = (Just 3, drop 5 str)
  | "four" `isPrefixOf` str = (Just 4, drop 4 str)
  | "five" `isPrefixOf` str = (Just 5, drop 4 str)
  | "six" `isPrefixOf` str = (Just 6, drop 3 str)
  | "seven" `isPrefixOf` str = (Just 7, drop 5 str)
  | "eight" `isPrefixOf` str = (Just 8, drop 5 str)
  | "nine" `isPrefixOf` str = (Just 9, drop 4 str)
  | otherwise = (Nothing, tail str)
