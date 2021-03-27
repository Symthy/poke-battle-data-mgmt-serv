import { Module } from '@nestjs/common';
import { PokemonsController } from './controller/pokemons.controller';
import { PokemonsService } from './service/pokemons.service';

@Module({
  imports: [],
  controllers: [PokemonsController],
  providers: [PokemonsService]
})
export class PokemonsModule { }

@Module({
  imports: [PokemonsModule],
})
export class AppModule { }
