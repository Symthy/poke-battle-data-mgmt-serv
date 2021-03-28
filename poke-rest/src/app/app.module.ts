import { Module } from '@nestjs/common';
import { AbilitiesController } from './controller/abilities.controller';
import { MovesController } from './controller/moves.controller';
import { PokemonsController } from './controller/pokemons.controller';
import { TypesController } from './controller/types.controller';
import { AbilitiesService } from './provider/abilities.service';
import { MovesService } from './provider/moves.service';
import { PokemonsService } from './provider/pokemons.service';
import { TypesService } from './provider/types.service';

@Module({
  imports: [],
  controllers: [PokemonsController],
  providers: [PokemonsService]
})
export class PokemonsModule { }

@Module({
  imports: [],
  controllers: [AbilitiesController],
  providers: [AbilitiesService]
})
export class AbilitiesModule { }

@Module({
  imports: [],
  controllers: [MovesController],
  providers: [MovesService]
})
export class MovesModule { }

@Module({
  imports: [],
  controllers: [TypesController],
  providers: [TypesService]
})
export class TypesModule { }

@Module({
  imports: [PokemonsModule, AbilitiesModule, MovesModule, TypesModule],
})
export class AppModule { }
