const fs = require("fs");
const https = require("https");

const token = process.env.GITHUB_TOKEN;
const aula = process.env.AULA;

if (!token) {
  console.error("Erro: GITHUB_TOKEN não definido.");
  process.exit(1);
}

if (!aula) {
  console.error("Erro: AULA não definida.");
  process.exit(1);
}

const reposFile = "repos.txt";

if (!fs.existsSync(reposFile)) {
  console.error("Arquivo repos.txt não encontrado.");
  process.exit(1);
}

const lines = fs.readFileSync(reposFile, "utf-8").split("\n").filter(Boolean);

console.log("Chamada:");

lines.forEach((line) => {
  const [name, repoURL] = line.split(" | ");

  const parts = repoURL.split("github.com/");
  if (parts.length < 2) {
    console.log(`${name} - Erro ao dividir URL do repositório`);
    return;
  }

  const repoPath = parts[1];
  const apiURL = `https://api.github.com/repos/${repoPath}/commits`;

  const options = {
    headers: {
      "User-Agent": "presenca-script",
      Authorization: `Bearer ${token}`,
    },
  };

  https.get(apiURL, options, (res) => {
    let data = "";

    res.on("data", (chunk) => {
      data += chunk;
    });

    res.on("end", () => {
        try {
        const commits = JSON.parse(data);
        if (!Array.isArray(commits)) {
            console.log(`${name}\t\t0`);
          return;
        }            
        const found = commits.some((commit) =>
          commit.commit.message.includes(`aula ${aula} -`)
        );

        console.log(`${name}\t\t${found ? 1 : 0}`);
      } catch (err) {
        // console.log(`${name} - Erro ao processar resposta da API: ${err.message}`);
      }
    });
  }).on("error", (err) => {
    console.log(`${name} - Erro na requisição`);
  });
});
