import pycantonese
import json
import sys

def convert_to_jyutping(text):
    try:
        result = pycantonese.characters_to_jyutping(text)
        return [{"character": char, "jyutping": jyut} for char, jyut in result]
    except Exception as e:
        return [{"character": text, "jyutping": "", "error": str(e)}]

if __name__ == "__main__":
    text = sys.argv[1] if len(sys.argv) > 1 else ""
    result = convert_to_jyutping(text)
    print(json.dumps(result, ensure_ascii=False))
