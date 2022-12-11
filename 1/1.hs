getsFloats :: FilePath -> [[Int]]
getsFloats path = do
  contents <- readFile path
  let ints = map read  . lines $ contents
  return ints

  --groupify = filter (/= [""]) . groupBy (\x y -> x /= "" && y /= "") . lines


main :: IO ()
main =
  do
    print (getsFloats)

