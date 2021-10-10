import { Pokemon, Pokemons } from '@/models/pokemons';
import { Injectable } from '@nestjs/common';

@Injectable()
export class PokemonsService {
  getPokemons(): Pokemons {
    return new Pokemons();
  }

  getPokemonById(id: string): Pokemon {
    return new Pokemon();
  }
}
