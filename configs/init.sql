INSERT INTO hospitals (id, name_th, name_en, created_at, updated_at)
VALUES (
    'eb83e5fb-2662-4edf-b7ec-4f7b8cd6feb5',
    'โรงพยาบาลสินแพทย์',
    'synphate hospital',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
),
(
    'af64cd87-8683-4c10-a60b-a5c2d7d9a725',
    'โรงพยาบาลรามา',
    'Rama hospital',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
),
(
    '49d50c07-4066-45b5-a788-15c8ce00128d',
    'โรงพยาบาลแพทย์รังสิต',
    'phate rangsit hospital',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
);

INSERT INTO patients (id, hospital_id, email, first_name_th, last_name_th, middle_name_th, first_name_en, last_name_en, middle_name_en, date_of_birth, patient_hn, national_id,
passport_id, gender, phone_number, country_code, created_at, updated_at)
VALUES (
    '29407b86-4a26-4937-b81c-4119dbceec42',
    'eb83e5fb-2662-4edf-b7ec-4f7b8cd6feb5',
    'patient1@example.com',
    'สมชาย',
    'สมบูรณ์',
    '-',
    'Somchai',
    'Sombut',
    '-',
    '1990-01-01',
    'HN123455',
    '1234567890',
    '134564662',
    'M',
    '8434567453',
    '+66',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
),
(
    '683b9b68-6476-4697-9f07-299efdaadd99',
    'af64cd87-8683-4c10-a60b-a5c2d7d9a725',
    'patient2@example.com',
    'สมหญิง',
    'สมบูรณ์',
    '-',
    'Somying',
    'Sombut',
    '-',
    '1990-01-03',
    'HN12345634',
    '123456789220',
    '134564663234',
    'F',
    '902345645',
    '+66',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
),
(
    '3da4b072-b36e-4100-8282-fa168f23149b',
    '49d50c07-4066-45b5-a788-15c8ce00128d',
    'patient3@example.com',
    'สมใจ',
    'สมบูรณ์',
    '-',
    'Somjai',
    'Somsong',
    '-',
    '1990-01-02',
    'HN757456',
    '33454664',
    '3456788954',
    'F',
    '609994444',
    '+66',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
);