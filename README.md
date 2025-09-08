# Hangulize Cantonese

한글라이즈 라이브러리에 광둥어(번체 중국어) 지원을 추가한 프로젝트입니다.

## 개요

이 프로젝트는 [hangulize](https://github.com/hangulize/hangulize) 라이브러리를 기반으로 하여, 광둥어를 LSHK(jyutping) 표기법으로 변환한 후 한글로 표기하는 기능을 제공합니다.

## 주요 기능

- **광둥어 LSHK 표기법 지원**: LSHK jyutping 입력을 한글로 변환
- **한자→jyutping 변환**: Python 라이브러리(pycantonese)를 통한 번체 한자 자동 변환
- **HSL 규칙 파일**: `specs/yue.hsl`에 광둥어 전사 규칙 정의
- **jyutping 변환 패키지**: `translit/jyutping/`에 LSHK 변환 로직 구현
- **포괄적인 테스트**: 다양한 광둥어 음운 환경에 대한 테스트 케이스 포함

## 설치 및 사용

### 필수 요구사항
- **Go 1.19 이상**
- **Python 3.7 이상** (한자→jyutping 변환용)
- **pycantonese 라이브러리**

### Python 의존성 설치
```bash
# pycantonese 설치 (한자→jyutping 변환용)
pip install pycantonese

# 또는 requirements.txt 사용
pip install -r requirements.txt
```

### Go 설치 및 빌드
```bash
# 프로젝트 클론
git clone https://github.com/jskim947/hangulize-cantonese.git
cd hangulize-cantonese

# 빌드
go build cmd/hangulize/main.go

# 실행
./main yue "nei5 hou2"        # 안녕하세요
./main yue "hoeng1 gong2"     # 홍콩
./main yue "ngo5"             # 나
```

### 테스트 실행
```bash
# LSHK jyutping 직접 입력 테스트
go run cmd/hangulize/main.go yue "nei5 hou2"        # 안녕하세요
go run cmd/hangulize/main.go yue "hoeng1 gong2"     # 홍콩
go run cmd/hangulize/main.go yue "ngo5"             # 나

# Python을 통한 한자→jyutping→한글 변환 테스트
python test_jyutping.py

# 또는 직접 Python에서 사용
python -c "
import pycantonese
import subprocess
import sys

# 한자를 jyutping으로 변환
text = '香港你好'
jyutping = pycantonese.characters_to_jyutping(text)
print(f'한자: {text}')

# jyutping을 한글로 변환
for word, jp in jyutping:
    result = subprocess.run(['go', 'run', 'cmd/hangulize/main.go', 'yue', jp], 
                          capture_output=True, text=True)
    hangul = result.stdout.strip()
    print(f'{word} ({jp}) → {hangul}')
"
```

## 지원하는 광둥어 음운

### 자음
- **기본 자음**: b, p, m, f, d, t, n, l, g, k, h, z, c, s, w
- **복합 자음**: gw, kw, ng
- **j 초성**: j + 모음 결합

### 모음
- **단일 모음**: a, e, i, o, u, aa, oe, yu
- **이중모음**: ai, au, ei, ou, ui, iu, aai, aau, eoi
- **삼중모음**: aai, aau, eoi
- **w 포함 모음**: wa, wai, wan, wong, wing 등
- **y 운모**: yu, yun, yut

### 성조
- **성조 1-6**: LSHK 표기법의 성조 번호는 한글 표기에 영향을 주지 않음

## 예시

| 한자 | LSHK jyutping | 한글 표기 | 의미 |
|------|---------------|-----------|------|
| 你好 | nei5 hou2 | 네이호우 | 안녕하세요 |
| 香港 | hoeng1 gong2 | 헝공 | 홍콩 |
| 我 | ngo5 | 응오 | 나 |
| 唔該 | m4 goi1 | 므고이 | 죄송합니다 |
| 廣東話 | gwong2 dung1 waa2 | 궝둥와 | 광둥어 |
| 多謝 | do1 ze6 | 도제 | 감사합니다 |
| 再見 | zoi3 gin3 | 조이긴 | 안녕히 가세요 |

## 프로젝트 구조

```
hangulize-cantonese/
├── specs/
│   └── yue.hsl              # 광둥어 HSL 규칙 파일
├── translit/
│   └── jyutping/
│       └── jyutping.go      # LSHK 변환 패키지
├── cmd/hangulize/
│   └── main.go              # CLI 도구
├── test_jyutping.py         # Python 테스트 스크립트
└── README.md               # 이 파일
```

## HSL 규칙 파일

`specs/yue.hsl` 파일에는 광둥어를 한글로 변환하는 모든 규칙이 정의되어 있습니다:

- **rewrite**: 성조 번호 제거 등 전처리 규칙
- **transcribe**: LSHK jyutping을 한글 자모로 변환하는 규칙
- **test**: 다양한 테스트 케이스

## 기여하기

1. 이 저장소를 포크합니다
2. 새로운 기능 브랜치를 생성합니다 (`git checkout -b feature/amazing-feature`)
3. 변경사항을 커밋합니다 (`git commit -m 'Add amazing feature'`)
4. 브랜치에 푸시합니다 (`git push origin feature/amazing-feature`)
5. Pull Request를 생성합니다

## 라이선스

이 프로젝트는 원본 hangulize 프로젝트와 동일한 라이선스를 따릅니다.

## 원본 프로젝트

- [hangulize](https://github.com/hangulize/hangulize) - 원본 한글라이즈 라이브러리
- [LSHK](https://www.lshk.org/jyutping) - 광둥어 로마자 표기법

## 개발자

- 프로젝트 기반: [hangulize](https://github.com/hangulize/hangulize)
- 광둥어 지원 추가: jskim

## 버전 히스토

- **v1.0.0** - 초기 릴리스
  - 광둥어 LSHK jyutping 지원
  - 기본 자음/모음 변환 규칙
  - 포괄적인 테스트 케이스