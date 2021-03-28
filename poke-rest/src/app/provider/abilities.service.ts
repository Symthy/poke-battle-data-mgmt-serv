import { Injectable } from '@nestjs/common';
import { Abilities, Ability } from '../models/abilities';

@Injectable()
export class AbilitiesService {
  getAbilityById(id: string): Ability {
    throw new Error('Method not implemented.');
  }

  getAbilities(): Abilities {
    return new Abilities();
  }
}
