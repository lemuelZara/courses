# Arrays 101

## Introdução

**Arrays** são estruturas de dados fundamentais na computação, que armazenam uma coleção de elementos (geralmente do mesmo tipo) em posições sequenciais de memória. Permitem acesso rápido a qualquer elemento através de seu índice, começando geralmente do zero.

Arrays são amplamente usados para representar dados de forma ordenada, como listas de notas, tabelas ou imagens em pixels.

Exemplo em Go:

```go
// Criando e usando um array em Go
var numeros [5]int     // Array de 5 inteiros
numeros[0] = 10        // Atribuição

fmt.Println(numeros[0]) // Acesso ao primeiro elemento (saída: 10)
```

### Entendendo Arrays em baixo nível (Low Level)

No nível mais baixo, arrays são blocos de memória contínua onde os elementos são armazenados um após o outro sem espaços entres eles:

- O endereço de um elemento `i` é dado por:
  - **Endereço base + `i` * Tamanho do elemento**
- Por isso, acessar qualquer índice é **O(1)**, pois basta calcular o endereço de memória.
- Arrays tem tamanho fixo após sua criação (na maioria das linguagens, incluindo Go).

Exemplo conceitual:

```diff
+------+------+------+------+
|  10  |  27  |  -5  |  3   |
+------+------+------+------+
   0      1      2      3   <- índices
```

Imagine um array de inteiros, começando no endereço de memória 1000, onde cada inteiro ocupa 4 bytes:

| Índice | Endereço (base + i×4) | Valor |
| ------ | --------------------- | ----- |
| 0      | 1000                  | 7     |
| 1      | 1004                  | 13    |
| 2      | 1008                  | 5     |
| 3      | 1012                  | -2    |

A localização de qualquer elemento é feita de forma direta, sem varredura ou busca linear.

## Resumo

- **Arrays** armazenam elementos do mesmo tipo em sequência.
- Permitem acesso constante a qualquer posição.
- O tamanho do array é fixo após a criação.
- Operações comuns: leitura, escrita, iteração.

## Conceitos básicos

1. **Declaração e Inicialização**
    ```go
    var a [5]int             // array de 5 inteiros, valores padrões (zero)
    b := [3]string{"a","b","c"} // inicialização com valores
    ```

2. **Acesso e Modificação**
    ```go
    a[2] = 100
    x := a[2]
    ```

3. **Comprimento**
    ```go
    fmt.Println(len(a))
    ```

4. **Iteração**
    ```go
    for i, v := range a {
        fmt.Println(i, v)
    }
    ```

## Complexidade das operações (Big O)

| Operação              | Complexidade (Big O) |
| --------------------- | -------------------- |
| Acesso (Leitura)      | O(1)                 |
| Atualização (Escrita) | O(1)                 |
| Busca                 | O(n)                 |
| Inserção no início    | O(n)                 |
| Inserção no fim       | O(1)*                |
| Remoção no início     | O(n)                 |
| Remoção no fim        | O(1)                 |

> **Em arrays fixos (como em Go), inserção/remoção não é possível sem criar um novo array.*

## Técnicas mais comuns associadas a Arrays

- **Busca Linear (Linear Search):** procurar um elemento sequencialmente.
- **Busca Binária (Binary Search):** procurar elemento em array ordenado (dividindo em duas partes).
- **Dois Ponteiros (Two Pointers):** percorrer o array de dois lados simultaneamente.
- **Sliding Window (Janela Deslizante):** manter uma janela fixa ou variável enquanto processa o array.
- **Prefix Sum (Soma Prefixada):** processar somas acumuladas para consultas rápidas.
- **Reverse (Inversão):** trocar elementos de lugar para inverter o array.
- **Rotação:** deslocar todos os elementos para a direita ou esquerda.

## Exercício clássicos

Abaixo estão exercícios populares, do mais fácil ao mais difícil, cobrindo várias técnicas:

| Exercício                | Nível  | O que explora                                                   | Link                                                                              |
| ------------------------ | ------ | --------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| Two Sum                  | Easy   | Array, Hash Table                                               | [LeetCode 1](https://leetcode.com/problems/two-sum/)                              |
| Remove Duplicates Sorted | Easy   | Array, Two Pointers                                             | [LeetCode 26](https://leetcode.com/problems/remove-duplicates-from-sorted-array/) |
| Contains Duplicate II    | Easy   | Array, Hash Table, Sliding Window                               | [LeetCode 219](https://leetcode.com/problems/contains-duplicate-ii)               |
| Move Zeroes              | Easy   | Array, Two Pointers                                             | [LeetCode 283](https://leetcode.com/problems/move-zeroes/)                        |
| Maximum Subarray         | Medium | Array, Divide and Conquer, Dynamic Programming                  | [LeetCode 53](https://leetcode.com/problems/maximum-subarray/)                    |
| Subarray Sum Equals K    | Medium | Array, Hash Table, Prefix Sum                                   | [LeetCode 560](https://leetcode.com/problems/subarray-sum-equals-k/)              |
| Trapping Rain Water      | Hard   | Array, Two Pointer, Dynamic Programming, Stack, Monotonic Stack | [LeetCode 42](https://leetcode.com/problems/trapping-rain-water/)                 |