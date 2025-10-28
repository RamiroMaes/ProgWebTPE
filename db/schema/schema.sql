DROP TABLE IF EXISTS Tiene CASCADE;
DROP TABLE IF EXISTS Jugo CASCADE;
DROP TABLE IF EXISTS Jugador CASCADE;
DROP TABLE IF EXISTS Club CASCADE;
DROP TABLE IF EXISTS Lesion CASCADE;
DROP TABLE IF EXISTS Pais CASCADE;

CREATE TABLE Club (
    Nombre varchar(40)  NOT NULL,
    Ciudad varchar(40)  NOT NULL,
    CONSTRAINT Club_pk PRIMARY KEY (Nombre,Ciudad)
);
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

-- Table: Jugo
CREATE TABLE Jugo (
    fecha_inicio date  NOT NULL,
    fecha_fin date  NOT NULL,
    Jugador_iD_Jugador int  NOT NULL,
    Club_Nombre varchar(40)  NOT NULL,
    Club_Ciudad varchar(40)  NOT NULL,
    CONSTRAINT Jugo_en PRIMARY KEY (Jugador_iD_Jugador,Club_Nombre,Club_Ciudad)
);

-- Table: Lesion
CREATE TABLE Lesion (
    Tipo_Lesion varchar(50)  NOT NULL,
    Descripcion varchar(100)  NOT NULL,
    CONSTRAINT Lesion_pk PRIMARY KEY (Tipo_Lesion)
);

-- Table: Pais
CREATE TABLE Pais (
    Nombre varchar(40)  NOT NULL,
    CONSTRAINT Pais_pk PRIMARY KEY (Nombre)
);

-- Table: Tiene
CREATE TABLE Tiene (
    fecha_inicio date  NOT NULL,
    fecha_fin date  NOT NULL,
    Jugador_iD_Jugador int  NOT NULL,
    Lesion_Tipo_Lesion varchar(50)  NOT NULL,
    CONSTRAINT Tiene_pk PRIMARY KEY (Jugador_iD_Jugador,Lesion_Tipo_Lesion)
);

-- foreign keys
-- Reference: Jugador_Pais (table: Jugador)
ALTER TABLE Jugador ADD CONSTRAINT Jugador_Pais
    FOREIGN KEY (Pais_Nombre)
    REFERENCES Pais (Nombre)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Jugo_Club (table: Jugo)
ALTER TABLE Jugo ADD CONSTRAINT Jugo_Club
    FOREIGN KEY (Club_Nombre, Club_Ciudad)
    REFERENCES Club (Nombre, Ciudad)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Jugo_Jugador (table: Jugo)
ALTER TABLE Jugo ADD CONSTRAINT Jugo_Jugador
    FOREIGN KEY (Jugador_iD_Jugador)
    REFERENCES Jugador (iD_Jugador)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Tiene_Jugador (table: Tiene)
ALTER TABLE Tiene ADD CONSTRAINT Tiene_Jugador
    FOREIGN KEY (Jugador_iD_Jugador)
    REFERENCES Jugador (iD_Jugador)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Tiene_Lesion (table: Tiene)
ALTER TABLE Tiene ADD CONSTRAINT Tiene_Lesion
    FOREIGN KEY (Lesion_Tipo_Lesion)
    REFERENCES Lesion (Tipo_Lesion)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- End of file.
