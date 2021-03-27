import { Injectable } from '@nestjs/common';
import { Pokemons } from 'src/controller/responseModel/pokemons';

@Injectable()
export class PokemonsService {
  getPokemons(): Pokemons {
    return new Pokemons();
  }
}
