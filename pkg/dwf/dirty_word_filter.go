/**
  @author: ZYL
  @date: 2022/5/2
  @note: 用于过滤/替换敏感词的包， 使用Tire算法实现
*/
package dwf

import (
	"bufio"
	"errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"io"
	"os"
	"unicode/utf8"
)

type Trie struct {
	child map[rune]*Trie
	word  string
}

var tire *Trie

func NewTrie() *Trie {
	return &Trie{
		child: make(map[rune]*Trie),
		word:  "",
	}
}
func (trie *Trie) insert(word string) *Trie {
	cur := trie
	for _, v := range []rune(word) {
		// 若存在，不做处理，若不存在，创建新的子树
		if _, ok := cur.child[v]; !ok {
			t := NewTrie()
			cur.child[v] = t
		}
		cur = cur.child[v]
	}
	cur.word = word
	return trie
}

// FilterString 过滤敏感词
func (trie *Trie) filterString(word string, replaceWord string) string {
	cur := trie
	for i, v := range []rune(word) {
		if _, ok := cur.child[v]; ok {
			cur = cur.child[v]
			if cur.word != "" {
				word = ReplaceStr(word, replaceWord, i+1-utf8.RuneCountInString(cur.word), i)
				cur = trie // ，符合条件，从头开始准备下一次遍历
			}
		} else {
			cur = trie // 不存在，则从头遍历
		}
	}
	return word
}

// ReplaceStr 替换敏感词
func ReplaceStr(word, replace string, left, right int) string {
	str := ""
	for i, v := range []rune(word) {
		if i >= left && i <= right {
			str = str + replace
		} else {
			str += string(v)
		}
	}
	return str
}

// FilterDirtyWord 对外暴露的过滤敏感词方法
func FilterDirtyWord(originStr string, replaceWord string) (error, string) {
	if tire == nil {
		return errors.New("Tire 未初始化"), originStr
	}

	return nil, tire.filterString(originStr, replaceWord)
}

func Init() error {
	tire = NewTrie()
	filePath := viper.GetString("dirtyword.path")
	file, err := os.Open(filePath)
	if err != nil {
		zap.L().Error("打开敏感词文件失败, err", zap.Error(err))
		return err
	}

	defer file.Close()
	reader := bufio.NewReader(file) // 读取文本数据
	for {
		word, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				zap.L().Error("读取文件失败， err=", zap.Error(err))
				return err
			}
		}
		wordStr := string(word)
		tire.insert(wordStr)
	}
	return nil
}

//func main() {
//	trie := NewTrie()
//	trie.insert("sb").insert("狗日").insert("cnm").insert("狗日的").insert("c").insert("nm")
//	fmt.Println(trie.filterString("狗头，你就是个狗日的，我要cnm，你个sb，嘿嘿"))
//}
