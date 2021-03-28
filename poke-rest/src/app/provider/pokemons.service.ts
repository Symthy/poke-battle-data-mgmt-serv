import { Injectable } from '@nestjs/common';
import { Pokemon, Pokemons } from 'src/app/models/pokemons';

@Injectable()
export class PokemonsService {
  getPokemons(): Pokemons {
    return new Pokemons();
  }

  getPokemonById(id: string): Pokemon {
    return new Pokemon();
  }
}
