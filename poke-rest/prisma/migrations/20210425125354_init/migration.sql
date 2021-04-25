-- CreateEnum
CREATE TYPE "Role" AS ENUM ('USER', 'ADMIN');

-- CreateEnum
CREATE TYPE "Species" AS ENUM ('Physical', 'Special');

-- CreateEnum
CREATE TYPE "Type" AS ENUM ('Normal', 'Fighting', 'Flying', 'Poison', 'Ground', 'Rock', 'Bug', 'Ghost', 'Steel', 'Fire', 'Water', 'Grass', 'Electric', 'Psychic', 'Ice', 'Dragon', 'Dark', 'None');

-- CreateTable
CREATE TABLE "profile" (
    "id" SERIAL NOT NULL,
    "bio" TEXT,
    "userId" INTEGER NOT NULL,

    PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "user" (
    "id" SERIAL NOT NULL,
    "email" TEXT NOT NULL,
    "name" VARCHAR(20) NOT NULL,
    "role" "Role" NOT NULL DEFAULT E'USER',

    PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "pokemons" (
    "id" INTEGER NOT NULL,
    "name" VARCHAR(20) NOT NULL,
    "type1" "Type" NOT NULL,
    "type2" "Type" NOT NULL,
    "base_stats_h" INTEGER NOT NULL DEFAULT 0,
    "base_stats_a" INTEGER NOT NULL DEFAULT 0,
    "base_stats_b" INTEGER NOT NULL DEFAULT 0,
    "base_stats_c" INTEGER NOT NULL DEFAULT 0,
    "base_stats_d" INTEGER NOT NULL DEFAULT 0,
    "base_stats_s" INTEGER NOT NULL DEFAULT 0,

    PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "abilities" (
    "id" SERIAL NOT NULL,
    "name" VARCHAR(20) NOT NULL,
    "description" TEXT,
    "physical_move_power_correction" INTEGER NOT NULL DEFAULT 1,
    "special_move_power_correction" INTEGER NOT NULL DEFAULT 1,
    "attack_power_correction" INTEGER NOT NULL DEFAULT 1,
    "special_attack_power_correction" INTEGER NOT NULL DEFAULT 1,
    "attack_correction" INTEGER NOT NULL DEFAULT 1,
    "special_attack_correction" INTEGER NOT NULL DEFAULT 1,
    "damage_correction" INTEGER NOT NULL DEFAULT 1,
    "weight_correction" INTEGER NOT NULL DEFAULT 1,

    PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "moves" (
    "id" SERIAL NOT NULL,
    "name" VARCHAR(20) NOT NULL,
    "type" "Type" NOT NULL,
    "move_species" "Species" NOT NULL,
    "powers" INTEGER NOT NULL,
    "accuracy" INTEGER NOT NULL,
    "pp" INTEGER NOT NULL,
    "is_contact_move" BOOLEAN NOT NULL,
    "is_can_guard" BOOLEAN NOT NULL,

    PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "type_compatibility" (
    "attack_type" "Type" NOT NULL,
    "defence_type" "Type" NOT NULL,
    "compatibility" INTEGER NOT NULL DEFAULT 1,

    PRIMARY KEY ("attack_type","defence_type")
);

-- CreateTable
CREATE TABLE "trained_pokemons" (
    "id" SERIAL NOT NULL,
    "effort_value_h" INTEGER NOT NULL,
    "effort_value_a" INTEGER NOT NULL,
    "effort_value_b" INTEGER NOT NULL,
    "effort_value_c" INTEGER NOT NULL,
    "effort_value_d" INTEGER NOT NULL,
    "effort_value_s" INTEGER NOT NULL,
    "pokemon_id" INTEGER NOT NULL,

    PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "party" (
    "user_id" INTEGER NOT NULL,
    "trained_pokemon_id" INTEGER NOT NULL,
    "move1" INTEGER NOT NULL,
    "move2" INTEGER NOT NULL,
    "move3" INTEGER NOT NULL,
    "move4" INTEGER NOT NULL,

    PRIMARY KEY ("user_id","trained_pokemon_id")
);

-- CreateIndex
CREATE UNIQUE INDEX "profile.userId_unique" ON "profile"("userId");

-- CreateIndex
CREATE UNIQUE INDEX "user.email_unique" ON "user"("email");

-- CreateIndex
CREATE UNIQUE INDEX "pokemons.name_unique" ON "pokemons"("name");

-- CreateIndex
CREATE UNIQUE INDEX "abilities.name_unique" ON "abilities"("name");

-- CreateIndex
CREATE UNIQUE INDEX "moves.name_unique" ON "moves"("name");

-- AddForeignKey
ALTER TABLE "profile" ADD FOREIGN KEY ("userId") REFERENCES "user"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "trained_pokemons" ADD FOREIGN KEY ("pokemon_id") REFERENCES "pokemons"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "party" ADD FOREIGN KEY ("user_id") REFERENCES "user"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "party" ADD FOREIGN KEY ("trained_pokemon_id") REFERENCES "trained_pokemons"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "party" ADD FOREIGN KEY ("move1") REFERENCES "moves"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "party" ADD FOREIGN KEY ("move2") REFERENCES "moves"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "party" ADD FOREIGN KEY ("move3") REFERENCES "moves"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "party" ADD FOREIGN KEY ("move4") REFERENCES "moves"("id") ON DELETE CASCADE ON UPDATE CASCADE;
