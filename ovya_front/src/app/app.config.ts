import { importProvidersFrom, LOCALE_ID } from '@angular/core';
import { provideHttpClient, withInterceptorsFromDi } from '@angular/common/http';
import { RouterModule } from '@angular/router';
import { ReactiveFormsModule, FormsModule } from '@angular/forms';

import { routes } from './app.routes';

export const appConfig = {
  providers: [
    provideHttpClient(withInterceptorsFromDi()),
    importProvidersFrom(
      ReactiveFormsModule,
      FormsModule,
      RouterModule.forRoot(routes)
    ),
    {
      provide: LOCALE_ID,
      useValue: 'fr-FR'
    }
  ]
};