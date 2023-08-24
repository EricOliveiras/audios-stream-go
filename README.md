# Audios stream Go

Este é um projeto simples em Golang que demonstra como criar um servidor de streaming de áudio básico.

## Requisitos

Certifique-se de ter o Go instalado em seu sistema. Você pode baixá-lo em: https://golang.org/dl/

## Como Executar

1. Clone este repositório para o seu sistema:

```bash
git clone https://github.com/EricOliveiras/audios-stream-go.git
```

2. Navegue até o diretório do projeto:

```bash
cd audios-stream-go
```

3. Você pode executar o servidor de duas formas:

```bash
go run cmd/audio-stream/main.go path/to/your/audio/file.mp3
```

ou criar o executavel da aplicação:

```bash
go build cmd/audio-stream/main.go
```

Será gerado um arquivo chamado 'main' na raiz do projeto. Execute-o:

```bash
./main path/to/your/audio/file.mp3
```

Substitua `path/to/your/audio/file.mp3` pelo caminho real para o arquivo de áudio que você deseja transmitir.

> Lembre-se sempre de passar o caminho do aúdio.

4. Abra um navegador e acesse: http://localhost:8080/musics
