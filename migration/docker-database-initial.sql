create table bairrosgoiania(
    id serial primary key,
    nome varchar,
    cidade varchar,
    Ano_Criacao INTEGER
);

INSERT INTO bairrosgoiania(nome, cidade, Ano_Criacao) VALUES
		('Setor União', 'Goiâni', 1960),
		('Vila Alpes', 'Goiânia', 1960),
		('Jardim America', 'Goiânia', 1952),
		('Novo Horizonte', 'Goiânia', 1975),
		('Jardim Europa', 'Goiânia', 1922),
		('Vila São joão', 'Goiânia', 1968),
		('Vila Nova Esperança', 'Goiânia', 1979);
