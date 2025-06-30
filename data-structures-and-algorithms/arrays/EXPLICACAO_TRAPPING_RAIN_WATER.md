# 🌧️ Trapping Rain Water

## O que este algoritmo faz? 🤔

Imagine que você tem uma série de **barras de diferentes alturas** lado a lado, como se fossem prédios ou muros. Quando chove, a água fica "presa" entre essas barras, formando pequenas piscininhas.

**O algoritmo calcula quanto de água da chuva consegue ficar acumulada entre essas barras.**

## Uma analogia simples 🏠

Pense assim: você tem uma fileira de caixas de diferentes alturas no seu quintal. Quando chove:
- A água só fica "presa" se tiver uma caixa mais alta de cada lado
- Se não tiver "paredes" altas o suficiente, a água escorre para fora

**Exemplo visual:**

Tomando como base o array de `[0,1,0,2,1,0,1,3,2,1,2,1]`

<img src="./rainwatertrap.png" />

## Como o algoritmo funciona? ⚙️

O algoritmo usa uma técnica muito esperta chamada **"dois ponteiros"**:

### 1. **Duas pessoas trabalhando juntas** 👥
- Uma pessoa começa na **esquerda** (início da fileira)
- Outra pessoa começa na **direita** (fim da fileira)
- Elas vão se aproximando, cada uma cuidando do seu lado

### 2. **Cada pessoa lembra da maior altura que já viu** 📏
- A pessoa da esquerda lembra: "qual foi a barra mais alta que eu já passei?"
- A pessoa da direita lembra: "qual foi a barra mais alta que eu já passei?"

### 3. **Sempre mexe quem está no lado mais baixo** ⚖️
- Se a barra atual da esquerda é menor que a da direita:
  - A pessoa da esquerda avança
- Senão:
  - A pessoa da direita avança

### 4. **Calculando a água** 💧
Quando alguém avança, pergunta:
- "A barra mais alta que eu já vi consegue segurar água aqui?"
- Se sim: `água = altura_máxima_vista - altura_atual`
- Se não: não acumula água (altura atual é maior)

## Por que esse método funciona? 🧠

A **genialidade** está em sempre mover o lado com a barra menor:

- **Se eu estou no lado mais baixo**: posso calcular a água com segurança, porque sei que do outro lado tem uma barra alta o suficiente para "segurar" a água
- **O lado mais alto**: sempre vai conseguir conter a água, então não preciso me preocupar com ele agora

## Exemplo passo a passo 👣

Vamos usar: `[0,1,0,2,1,0,1,3,2,1,2,1]`

```
Início:
esquerda=0 (altura=0), direita=11 (altura=1)
maxEsquerda=0, maxDireita=0

Passo 1: altura[0]=0 < altura[11]=1
- maxEsquerda = max(0,0) = 0
- água += max(0-0, 0) = 0
- esquerda++

Passo 2: altura[1]=1 = altura[11]=1
- Como são iguais, mexe a direita
- maxDireita = max(0,1) = 1
- água += max(1-1, 0) = 0
- direita--

... e assim por diante até se encontrarem
```

## Vantagens deste algoritmo ✨

1. **Rápido**: Passa por cada barra apenas uma vez - O(n)
2. **Econômico**: Usa pouca memória extra - O(1)
3. **Elegante**: Resolve um problema aparentemente complicado de forma simples

## Analogia final para fixar 🎯

É como se duas pessoas estivessem **limpando uma calha** de chuva:
- Cada uma começa de um lado
- Sempre quem está na parte mais baixa avança
- Vão calculando quanta água conseguem "segurar" em cada ponto
- Se encontram no meio, terminaram o trabalho!

---

**Resumo**: Este algoritmo é um exemplo brilhante de como usar **lógica simples** e **organização esperta** (dois ponteiros) para resolver um problema que parece complicado de forma muito eficiente! 🚀
