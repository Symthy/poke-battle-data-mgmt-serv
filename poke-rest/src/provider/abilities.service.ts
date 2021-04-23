import { Abilities, Ability } from '@/models/abilities';
import { Injectable } from '@nestjs/common';

@Injectable()
export class AbilitiesService {
  getAbilityById(id: string): Ability {
    throw new Error('Method not implemented.');
  }

  getAbilities(): Abilities {
    return new Abilities();
  }
}
