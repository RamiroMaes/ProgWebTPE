DROP TABLE IF EXISTS Jugador CASCADE;
DROP TABLE IF EXISTS Pais CASCADE;

-- Table: Jugador
CREATE TABLE Jugador (
    Nombre varchar(40)  NOT NULL,
    iD_Jugador int  NOT NULL,
    Posicion varchar(40)  NOT NULL,
    fecha_nacimiento date  NOT NULL,
    altura int  NOT NULL,
    Pais_Nombre varchar(40)  NOT NULL,
    CONSTRAINT Jugador_pk PRIMARY KEY (iD_Jugador)
);


-- Table: Pais
CREATE TABLE Pais (
    Nombre varchar(40)  NOT NULL,
    CONSTRAINT Pais_pk PRIMARY KEY (Nombre)
);

-- foreign keys
-- Reference: Jugador_Pais (table: Jugador)
ALTER TABLE Jugador ADD CONSTRAINT Jugador_Pais
    FOREIGN KEY (Pais_Nombre)
    REFERENCES Pais (Nombre)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- End of file.